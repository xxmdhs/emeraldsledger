package thread

import (
	"bytes"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/xxmdhs/emeraldsledger/http"
)

var ErrNotFind = fmt.Errorf("not find")

var (
	//pagePidReg = regexp.MustCompile(`<div id="post_(\d{1,15}?)".*?>`)
	//postReg    = regexp.MustCompile(`<div id="post_\d{1,20}".*?>[\s\S]*?<tr class="ad">`)
	//uidReg     = regexp.MustCompile(`<div class="authi"><a href=".*?(\d{1,15})" target="_blank" class="xw1".*?>(.{1,15}?)</a>`)
	//msgReg     = regexp.MustCompile(`<td[^>]*id="postmessage_\d{1,20}"[^>]*>\n*([\s\S]*?)<div id="comment_\d{1,20}" class="cm">`)
	titleReg = regexp.MustCompile(`<title>(.*)</title>`)
	//lockReg    = regexp.MustCompile(`<div class="locked">.*?</div>`)
	//buttonReg  = regexp.MustCompile(`<button[^>]*>.*?</button>`)
	//scriptReg  = regexp.MustCompile(`<script[^>]*>.*?</script>`)
)

func getPagePid(tid, page int, cookie string, LimitGet *http.LimitGet) ([]postData, string, error) {
	b, err := LimitGet.Get(`https://www.mcbbs.net/forum.php?mod=viewthread&tid=`+strconv.Itoa(tid)+`&page=`+strconv.Itoa(page)+`&ordertype=1`, cookie)
	if err != nil {
		return nil, "", fmt.Errorf("GetPagePid: %w", err)
	}
	tl := titleReg.FindSubmatch(b)
	if tl == nil {
		return nil, "", ErrNotFind
	}
	title := string(tl[1])

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(b))
	if err != nil {
		return nil, "", fmt.Errorf("GetPagePid: %w", err)
	}
	tDom := doc.Find("table[id^=pid]")

	pl := make([]postData, 0, len(tDom.Nodes))

	for i := range tDom.Nodes {
		p := postData{}
		v := tDom.Eq(i)
		nameDom := v.Find("div.authi > a.xw1").First()
		p.Username = nameDom.Text()
		uid, _ := getFromUrl(nameDom, "href", "uid")
		p.Authorid = uid
		pid, has := v.Attr("id")
		if !has {
			return nil, "", fmt.Errorf("GetPagePid: %w", ErrNotFindAttr{Attr: "id"})
		}
		p.Pid, err = strconv.Atoi(strings.TrimPrefix(pid, "pid"))
		if err != nil {
			return nil, "", fmt.Errorf("GetPagePid: %w", err)
		}
		postDom := v.Find("td[id^=postmessage_]").First()
		postDom.Find(".tip,.attach_tips").Each(func(i int, s *goquery.Selection) {
			s.Remove()
		})
		h, err := postDom.Html()
		if err != nil {
			return nil, "", fmt.Errorf("GetPagePid: %w", err)
		}
		p.Message = h
		pl = append(pl, p)
	}
	if len(pl) == 0 {
		return nil, "", fmt.Errorf("GetPagePid: %w", ErrListNull)
	}
	return pl, title, nil
}

func getFromUrl(s *goquery.Selection, attr, q string) (string, error) {
	href, has := s.Attr(attr)
	if !has {
		return "", fmt.Errorf("getFromUrl attr: %v: %w", attr, ErrNotFindAttr{Attr: attr})
	}
	u, err := url.Parse(href)
	if err != nil {
		return "", fmt.Errorf("getFromUrl attr: %v: %w", attr, err)
	}
	return u.Query().Get(q), nil
}

type ErrNotFindAttr struct {
	Attr string
}

func (e ErrNotFindAttr) Error() string {
	return "找不到 " + e.Attr
}

var ErrListNull = errors.New("list 为空")
