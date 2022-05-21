package thread

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/xxmdhs/emeraldsledger/http"
	"github.com/xxmdhs/emeraldsledger/structs"
)

func FindPage(tid, page int, cookie string, LimitGet *http.LimitGet) ([]structs.McbbsAd, error) {
	stid := strconv.Itoa(tid)

	l, err := getPagePid(tid, page, "", LimitGet)
	if err != nil {
		return nil, fmt.Errorf("FindPage: %w", err)
	}

	adCh := make(chan structs.McbbsAd, 20)
	eCh := make(chan error, 20)

	w := sync.WaitGroup{}

	cxt, c := context.WithCancel(context.Background())

	for _, v := range l {
		w.Add(1)
		v := v
		go func() {
			defer w.Done()
			b, err := LimitGet.Get(`https://www.mcbbs.net/forum.php?mod=misc&action=viewratings&tid=`+stid+`&pid=`+strconv.Itoa(v.Pid)+`&inajax=1`, "")
			if err != nil {
				select {
				case eCh <- fmt.Errorf("FindPage: %w", err):
				case <-cxt.Done():
				}
				return
			}
			pl := getpinfen(b)
			for _, vv := range pl {
				if vv.Type == "宝石" {
					if vv.Num > 0 {
						continue
					}
					w.Add(1)
					select {
					case <-cxt.Done():
						w.Done()
						return
					case adCh <- structs.McbbsAd{
						Uid:      v.Authorid,
						Username: v.Username,
						Count:    vv.Num,
						Time:     vv.Time,
						Cause:    v.Message,
						Type:     stid,
						Link:     "https://www.mcbbs.net/forum.php?mod=redirect&goto=findpost&ptid=" + stid + "&pid=" + strconv.Itoa(v.Pid),
					}:
					}
				}
			}
		}()

	}

	go func() {
		w.Wait()
		c()
	}()

	ads := make([]structs.McbbsAd, 0, len(l))

	for {
		select {
		case ad := <-adCh:
			ads = append(ads, ad)
			w.Done()
		case err := <-eCh:
			c()
			return nil, err
		case <-cxt.Done():
			return ads, nil
		}
	}
}

var ErrNotFind = fmt.Errorf("not find")

var (
	pagePidReg = regexp.MustCompile(`<div id="post_(\d{1,15}?)".*?>`)
	postReg    = regexp.MustCompile(`<div id="post_\d{1,20}".*?>[\s\S]*?<tr class="ad">`)
	uidReg     = regexp.MustCompile(`<div class="authi"><a href=".*?(\d{1,15})" target="_blank" class="xw1".*?>(.{1,15}?)</a>`)
	msgReg     = regexp.MustCompile(`<td class="t_f" id="postmessage_\d{1,20}".*?>\n*([\s\S]*?)<div id="comment_\d{1,20}" class="cm">`)
)

func getPagePid(tid, page int, cookie string, LimitGet *http.LimitGet) ([]postData, error) {
	b, err := LimitGet.Get(`https://www.mcbbs.net/forum.php?mod=viewthread&tid=`+strconv.Itoa(tid)+`&page=`+strconv.Itoa(page)+`&ordertype=1`, cookie)
	if err != nil {
		return nil, fmt.Errorf("getPagePid: %w", err)
	}
	l := postReg.FindAll(b, -1)
	if l == nil {
		return nil, ErrNotFind
	}
	pl := make([]postData, 0, len(l))

	for _, v := range l {
		tpid := pagePidReg.FindSubmatch(v)
		if tpid == nil {
			continue
		}
		pid, err := strconv.Atoi(string(tpid[1]))
		if err != nil {
			continue
		}
		tuid := uidReg.FindSubmatch(v)
		if tuid == nil {
			continue
		}
		uid := string(tuid[1])
		name := string(tuid[2])

		tmsg := msgReg.FindSubmatch(v)
		if tmsg == nil {
			continue
		}
		msg := string(tmsg[1])
		msg = strings.TrimSpace(msg)

		pl = append(pl, postData{
			Pid:      pid,
			Authorid: uid,
			Username: name,
			Message:  msg,
		})

	}
	return pl, nil
}

type postData struct {
	Authorid string `json:"authorid"`
	Username string `json:"username"`
	Message  string `json:"message"`
	Pid      int
}
