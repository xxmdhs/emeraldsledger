package thread

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var c = http.Client{Timeout: 30 * time.Second}

func HttpGet(url, cookie string) ([]byte, error) {
	reqs, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("HttpGet: %w", err)
	}
	reqs.Header.Set("Accept", "*/*")
	reqs.Header.Add("accept-encoding", "gzip")
	reqs.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36")
	reqs.Header.Set("Cookie", cookie)
	rep, err := c.Do(reqs)
	if rep != nil {
		defer rep.Body.Close()
	}
	if err != nil {
		return nil, fmt.Errorf("HttpGet: %w", err)
	}
	if rep.StatusCode != http.StatusOK {
		return nil, Errpget{Msg: rep.Status, url: url}
	}
	var reader io.ReadCloser
	switch rep.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(rep.Body)
		if err != nil {
			return nil, fmt.Errorf("httpget: %w", err)
		}
		defer reader.Close()
	default:
		reader = rep.Body
	}
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("HttpGet: %w", err)
	}
	return b, nil
}

type Errpget struct {
	Msg string
	url string
}

func (h Errpget) Error() string {
	return "not 200: " + h.Msg + " " + h.url
}
