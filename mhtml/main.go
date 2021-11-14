package mhtml

import (
	"embed"
	"encoding/json"
	"html/template"
	"os"
	"sort"
	"strconv"
	templateTet "text/template"
	"time"

	"github.com/xxmdhs/emeraldsledger/structs"
)

//go:embed template
var f embed.FS

var t *template.Template

var svgTep *templateTet.Template

func init() {
	var err error
	t, err = template.ParseFS(f, "template/*.html")
	if err != nil {
		panic(err)
	}
	svgTep, err = templateTet.ParseFS(f, "template/*.svg")
	if err != nil {
		panic(err)
	}
}

type table struct {
	Title string
	List  []tableList
}

type tableList struct {
	Rank string
	Uid  string
	Name string
	Num  string
	Link template.HTML
}

func Make(b []byte) {
	m := map[string]structs.McbbsAd{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		panic(err)
	}
	all(m)

	day30 := int64((time.Hour * 24 * 30).Seconds())
	//day90 := int64((time.Hour * 24 * 90).Seconds())
	//day365 := int64((time.Hour * 24 * 365).Seconds())

	//tableHtml(m, "table30.html", "一月内绿宝石使用排行", day30)
	//tableHtml(m, "table90.html", "三月内绿宝石使用排行", day90)
	//tableHtml(m, "table365.html", "一年内绿宝石使用排行", day365)
	//tableHtml(m, "all.html", "总绿宝石使用排行", 0)
	tableSvg(m, "table30.svg", "一月内绿宝石使用排行", day30)
}

func tableHtml(m map[string]structs.McbbsAd, filename, Title string, btime int64) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = t.ExecuteTemplate(f, "table", table{Title: Title, List: makeTable(m, btime)})
	if err != nil {
		panic(err)
	}
}

func tableSvg(m map[string]structs.McbbsAd, filename, Title string, btime int64) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = svgTep.ExecuteTemplate(f, "table.svg", table{Title: Title, List: makeTable(m, btime)})
	if err != nil {
		panic(err)
	}
}

func all(m map[string]structs.McbbsAd) {
	all := 0
	for _, v := range m {
		all = all + v.Count
	}
	type temp struct {
		Title       string
		AllEmeralds int
	}
	//f, err := os.Create("index.html")
	//if err != nil {
	//	panic(err)
	//}
	//defer f.Close()
	//err = t.ExecuteTemplate(f, "index", temp{Title: "绿宝石", AllEmeralds: all})
	//if err != nil {
	//	panic(err)
	//}
	ff, err := os.Create("count.svg")
	if err != nil {
		panic(err)
	}
	defer ff.Close()
	err = svgTep.ExecuteTemplate(ff, "count.svg", temp{AllEmeralds: all})
	if err != nil {
		panic(err)
	}
}

func makeTable(m map[string]structs.McbbsAd, btime int64) []tableList {
	tl := []tableList{}
	templ := []structs.McbbsAd{}
	tempm := map[string]structs.McbbsAd{}
	atime := time.Now().Unix()
	for _, v := range m {
		if v.Time > atime-btime || btime == 0 {
			if vv, ok := tempm[v.Uid]; ok {
				v.Count = v.Count + vv.Count
			}
			tempm[v.Uid] = v
		}
	}
	for _, v := range tempm {
		templ = append(templ, v)
	}
	sort.Slice(templ, func(i, j int) bool {
		return templ[i].Count < templ[j].Count
	})
	for i, v := range templ {
		tl = append(tl, tableList{
			Rank: strconv.Itoa(i + 1),
			Uid:  v.Uid,
			Name: v.Username,
			Num:  strconv.Itoa(v.Count),
			Link: template.HTML(`<a href="./user.html?uid=` + v.Uid + `">link</a>`),
		})
	}
	return tl
}
