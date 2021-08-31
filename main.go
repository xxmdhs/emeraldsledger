package main

import (
	"encoding/json"
	"flag"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/xxmdhs/emeraldsledger/http"
	"github.com/xxmdhs/emeraldsledger/mcbbsad"
	"github.com/xxmdhs/emeraldsledger/structs"
	"github.com/xxmdhs/emeraldsledger/thread"
)

func main() {
	w := sync.WaitGroup{}
	adl := []structs.McbbsAd{}
	lock := sync.Mutex{}

	LimitGet := http.NewLimitGet(threadInt, sleepTime, retry)

	w.Add(1)
	go func() {
		for i := 0; i < c.Page["adPage"]; i++ {
			l, err := mcbbsad.FindPage(i, cookie, LimitGet)
			if err != nil {
				panic(err)
			}
			lock.Lock()
			adl = append(adl, l...)
			lock.Unlock()
			time.Sleep(10 * time.Second)
		}
		w.Done()
	}()

	for k, v := range c.Page {
		if k == "adPage" {
			continue
		}
		i, err := strconv.Atoi(k)
		if err != nil {
			panic(err)
		}
		v := v
		w.Add(1)
		go func() {
			ad := threadFind(i, v, LimitGet)
			lock.Lock()
			adl = append(adl, ad...)
			lock.Unlock()
			w.Done()
		}()
	}

	w.Wait()
	m := map[string]structs.McbbsAd{}
	for _, v := range adl {
		m[v.Hash()] = v
	}

	f, err := os.Create("data.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	en := json.NewEncoder(f)
	en.SetEscapeHTML(false)
	en.SetIndent("", "    ")
	err = en.Encode(m)
	if err != nil {
		panic(err)
	}
}

func threadFind(tid, page int, LimitGet *http.LimitGet) []structs.McbbsAd {
	adl := []structs.McbbsAd{}
	l := sync.Mutex{}
	w := sync.WaitGroup{}

	a := 0
	for i := 0; i < page; i++ {
		w.Add(1)
		go func() {
			a++
			ad, err := thread.FindPage(tid, i, LimitGet)
			if err != nil {
				panic(err)
			}
			l.Lock()
			adl = append(adl, ad...)
			l.Unlock()
			w.Done()
		}()
		if a > threadInt {
			w.Wait()
			a = 0
		}
	}
	w.Wait()

	return adl
}

var (
	threadInt int
	retry     int
	cookie    string
	sleepTime int
	c         conifg
)

type conifg struct {
	Page map[string]int
}

func init() {
	flag.IntVar(&threadInt, "thread", 8, "thread")
	flag.IntVar(&retry, "retry", 10, "retry")
	flag.IntVar(&sleepTime, "sleep", 500, "sleep")
	flag.Parse()
	cookie = os.Getenv("cookie")

	b, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &c)
	if err != nil {
		panic(err)
	}
}
