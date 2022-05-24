package thread

import (
	"fmt"
	"regexp"
	"strings"
)

func ConvertHtml(message string) (string, []string) {
	list, message := replaceimg(message)
	message = replaceurl(message)
	message = strings.TrimSuffix(message, "</td></tr></table>\n\n\n</div>")
	message = strings.TrimSpace(message)
	return message, list
}

func replaceimg(msg string) ([]string, string) {
	list := []string{}
	msg = regmcbbsimg.ReplaceAllStringFunc(msg, func(s string) string {
		l := regmcbbsimg.FindStringSubmatch(s)
		if len(l) != 2 {
			return ""
		}
		list = append(list, l[1])
		return fmt.Sprintf(`<img src="%s">`, l[1])
	})
	msg = regimg.ReplaceAllStringFunc(msg, func(s string) string {
		l := regimg.FindStringSubmatch(s)
		if len(l) != 2 {
			return ""
		}
		if strings.Contains(l[1], `file="`) {
			return s
		}
		list = append(list, l[1])
		return fmt.Sprintf(`<img src="%s">`, l[1])
	})
	msg = aimgReg.ReplaceAllStringFunc(msg, func(s string) string {
		l := aimgReg.FindStringSubmatch(s)
		if len(l) != 2 {
			return ""
		}
		list = append(list, l[1])
		return fmt.Sprintf(`<img src="%s">`, l[1])
	})
	return list, msg
}

func replaceurl(msg string) string {
	msg = urlreg.ReplaceAllStringFunc(msg, func(s string) string {
		l := urlreg.FindStringSubmatch(s)
		if len(l) != 3 {
			return ""
		}
		if l[1] == l[2] {
			return l[1]
		}
		return fmt.Sprintf(`<a href="%s" target="_blank">%s</a>`, l[1], l[2])
	})
	return msg
}

var (
	regimg      = regexp.MustCompile(`<img[^>]*src=['"]([^'"]+)[^>]*>`)
	regmcbbsimg = regexp.MustCompile(`<img[^>]*file=['"]([^'"]+)[^>]*>`)
	urlreg      = regexp.MustCompile(`<a[^>]*href=['"]([^'"]+)[^>]*?>(.*?)</a>`)
	aimgReg     = regexp.MustCompile(`\[aimg=(.*?)\]`)
)
