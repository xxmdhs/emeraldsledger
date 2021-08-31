package thread

import (
	"encoding/json"
	"fmt"
	"strconv"

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
	ads := []structs.McbbsAd{}
	for _, v := range t.Variables.Postlist {
		if v.Tid != "" && v.Pid != "" {
			b, err := LimitGet.Get(`https://www.mcbbs.net/forum.php?mod=misc&action=viewratings&tid=`+v.Tid+`&pid=`+v.Pid+`&inajax=1`, "")
			if err != nil {
				return nil, fmt.Errorf("FindPage: %w", err)
			}
			pl := getpinfen(b)
			for _, vv := range pl {
				if vv.Type == "宝石" {
					if vv.Num > 0 {
						continue
					}
					ads = append(ads, structs.McbbsAd{
						Uid:      v.Authorid,
						Username: v.Username,
						Count:    vv.Num,
						Time:     vv.Time,
						Cause:    v.Message,
						Type:     stid,
						Link:     "https://www.mcbbs.net/forum.php?mod=redirect&goto=findpost&ptid=" + stid + "&pid=" + v.Pid,
					})
				}
			}
		}
	}
	return ads, nil
}
