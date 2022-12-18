package thread

import (
	"context"
	"fmt"
	"strconv"
	"sync"

	"github.com/xxmdhs/emeraldsledger/http"
	"github.com/xxmdhs/emeraldsledger/structs"
)

func FindPage(tid, page int, cookie string, LimitGet *http.LimitGet) ([]structs.McbbsAd, error) {
	stid := strconv.Itoa(tid)
	l, _, err := getPagePid(tid, page, cookie, LimitGet)
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
			b, err := LimitGet.Get(`https://www.mcbbs.net/forum.php?mod=misc&action=viewratings&tid=`+stid+`&pid=`+strconv.Itoa(v.Pid)+`&inajax=1`, cookie)
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

type postData struct {
	Authorid string `json:"authorid"`
	Username string `json:"username"`
	Message  string `json:"message"`
	Pid      int
}
