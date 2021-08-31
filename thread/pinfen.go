package thread

import (
	"regexp"
	"strconv"
	"time"

	"github.com/xxmdhs/emeraldsledger/structs"
)

func getpinfen(data []byte) []pinfen {
	ps := []pinfen{}
	l := pinfenReg.FindAllSubmatch(data, -1)
	for _, v := range l {
		if len(v) == 7 {
			num, _ := strconv.Atoi(string(v[2]))
			t, err := time.ParseInLocation("2006-1-2 15:04", string(v[5]), structs.Shanhai)
			if err != nil {
				continue
			}
			p := pinfen{
				Type: string(v[1]),
				Num:  num,
				Uid:  string(v[3]),
				Name: string(v[4]),
				Text: string(v[6]),
				Time: t.Unix(),
			}
			ps = append(ps, p)
		}
	}
	return ps
}

type pinfen struct {
	Uid  string
	Name string
	Type string
	Text string
	Time int64
	Num  int
}

var pinfenReg = regexp.MustCompile(`<tr>\n<td>(.{1,}?) ([+-]\d{1,}) .{1,}?</td>\n<td><a href="home.php\?mod=space&amp;uid=(\d{1,20}?)">(.{1,30}?)</a></td>\n<td>.*(\d{4}-\d{1,2}-\d{1,2} \d{1,2}:\d{1,2}).*</td>\n<td>(.*)</td>\n</tr>`)
