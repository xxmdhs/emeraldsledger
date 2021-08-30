package mcbbsad

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"

	_ "time/tzdata"

	"github.com/xxmdhs/emeraldsledger/http"
	"github.com/xxmdhs/emeraldsledger/structs"
)

var reg = regexp.MustCompile(`<td><a href="home.php\?mod=space&amp;uid=(\d{1,})" target="_blank">(.{1,}?)</a></td>\n.*<td>.*</td>\n.*<td>(.*?)</td>\n.*<td>(\d*?)</td>\n.*<td>(.{1,})</td>`)

func FindPage(page, retry int, cookie string) ([]structs.McbbsAd, error) {
	b, err := http.RetryGet(`https://www.mcbbs.net/plugin.php?id=mcbbs_ad:ad_history&all=1&page=`+strconv.Itoa(page), cookie, retry)
	if err != nil {
		return nil, fmt.Errorf("findPage: %w", err)
	}
	l := reg.FindAllSubmatch(b, -1)
	ads := []structs.McbbsAd{}
	for _, v := range l {
		if len(v) != 6 {
			continue
		}
		ad := structs.McbbsAd{}
		ad.Uid = string(v[1])
		ad.Username = string(v[2])
		ad.Cause = string(v[3])
		ad.Count, _ = strconv.Atoi(string(v[4]))
		ad.Type = "mcbbsAd"

		//2021-8-28 09:47
		t, err := time.ParseInLocation("2006-1-2 15:04", string(v[5]), structs.Shanhai)
		if err != nil {
			return nil, fmt.Errorf("findPage: %w", err)
		}
		ad.Time = t.Unix()
		ads = append(ads, ad)
	}
	return ads, nil
}

var ErrNotFind = errors.New("not find")
