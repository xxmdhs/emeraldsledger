package http

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var c = http.Client{
	Timeout: time.Second * 60,
}

func Httpget(url string, cookie string) ([]byte, error) {
	reqs, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Httpget: %w", err)
	}
	reqs.Header.Set("Accept", "*/*")
	reqs.Header.Add("accept-encoding", "gzip")
	reqs.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36")
	if cookie != "" {
		reqs.Header.Set("Cookie", cookie)
	}
	rep, err := c.Do(reqs)
	if rep != nil {
		defer rep.Body.Close()
	}
	if err != nil {
		return nil, fmt.Errorf("Httpget: %w", err)
	}
	if rep.StatusCode != 200 {
		return nil, fmt.Errorf("Httpget: %w", Not200{code: rep.Status, url: url})
	}
	var reader io.ReadCloser
	switch rep.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(rep.Body)
		if err != nil {
			return nil, fmt.Errorf("Httpget: %w", err)
		}
		defer reader.Close()
	default:
		reader = rep.Body
	}
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("Httpget: %w", err)
	}
	return b, nil
}

type Not200 struct {
	code string
	url  string
}

func (n Not200) Error() string {
	return "Not200: " + n.code + " " + n.url
}

func RetryGet(url string, cookie string, i int) ([]byte, error) {
	var err error
	var b []byte
	for a := 0; a < i; a++ {
		b, err = Httpget(url, cookie)
		if err == nil {
			break
		}
		log.Println(err)
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		return nil, fmt.Errorf("RetryGet: %w", err)
	}
	return b, nil
}

type LimitGet struct {
	ch        chan struct{}
	sleepTime int
	retry     int
}

func NewLimitGet(limit, sleepTime, retry int) *LimitGet {
	return &LimitGet{
		ch:        make(chan struct{}, limit),
		sleepTime: sleepTime,
		retry:     retry,
	}
}

func (l *LimitGet) Get(url string, cookie string) ([]byte, error) {
	l.ch <- struct{}{}
	defer func() {
		<-l.ch
	}()
	b, err := RetryGet(url, cookie, l.retry)
	if err != nil {
		return nil, fmt.Errorf("LimitGet: %w", err)
	}
	time.Sleep(time.Duration(l.sleepTime) * time.Millisecond)
	return b, nil
}
