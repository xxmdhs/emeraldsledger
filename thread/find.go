package thread

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"

	"github.com/xxmdhs/emeraldsledger/http"
	"github.com/xxmdhs/emeraldsledger/structs"
)

func FindPage(tid, page int, LimitGet *http.LimitGet) ([]structs.McbbsAd, error) {
	stid := strconv.Itoa(tid)
	b, err := LimitGet.Get(`https://www.mcbbs.net/api/mobile/index.php?version=4&module=viewthread&tid=`+stid+`&page=`+strconv.Itoa(page)+"&extra=&ordertype=1", "")
	if err != nil {
		return nil, fmt.Errorf("FindPage: %w", err)
	}
	t := thread{}
	err = json.Unmarshal(b, &t)
	if err != nil {
		return nil, fmt.Errorf("FindPage: %w", err)
	}
	adCh := make(chan structs.McbbsAd, 30)
	eCh := make(chan error, len(t.Variables.Postlist))

	w := sync.WaitGroup{}

	for _, v := range t.Variables.Postlist {
		if v.Tid != "" && v.Pid != "" {
			w.Add(1)
			go func() {
				defer w.Done()
				b, err := LimitGet.Get(`https://www.mcbbs.net/forum.php?mod=misc&action=viewratings&tid=`+v.Tid+`&pid=`+v.Pid+`&inajax=1`, "")
				if err != nil {
					eCh <- fmt.Errorf("FindPage: %w", err)
				}
				pl := getpinfen(b)
				for _, vv := range pl {
					if vv.Type == "宝石" {
						if vv.Num > 0 {
							continue
						}
						w.Add(1)
						adCh <- structs.McbbsAd{
							Uid:      v.Authorid,
							Username: v.Username,
							Count:    vv.Num,
							Time:     vv.Time,
							Cause:    v.Message,
							Type:     stid,
							Link:     "https://www.mcbbs.net/forum.php?mod=redirect&goto=findpost&ptid=" + stid + "&pid=" + v.Pid,
						}
					}
				}
			}()
		}
	}
	cxt, c := context.WithCancel(context.Background())

	go func() {
		w.Wait()
		c()
	}()

	ads := make([]structs.McbbsAd, 0, len(t.Variables.Postlist))

	for {
		select {
		case ad := <-adCh:
			ads = append(ads, ad)
			w.Done()
		case err := <-eCh:
			return nil, err
		case <-cxt.Done():
			return ads, nil
		}
	}
}
