package main

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"io"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"

	"github.com/xxmdhs/emeraldsledger/http"
	"github.com/xxmdhs/emeraldsledger/mcbbsad"
	"github.com/xxmdhs/emeraldsledger/structs"
	"github.com/xxmdhs/emeraldsledger/thread"
)

func main() {
	w := sync.WaitGroup{}
	adl := make(chan structs.McbbsAd, 20)
	cxt, cancel := context.WithCancel(context.Background())

	LimitGet := http.NewLimitGet(threadInt, sleepTime, retry)

	go func() {
		f, err := os.Create("data.txt")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		bw := bufio.NewWriter(f)
		defer bw.Flush()
		for {
			select {
			case <-cxt.Done():
				return
			case ad := <-adl:
				b, err := json.Marshal(ad)
				if err != nil {
					panic(err)
				}
				bw.Write(b)
				bw.Write([]byte("\n"))
				log.Println(string(b))
			}
		}
	}()
	cc := make(chan os.Signal, 1)
	signal.Notify(cc, os.Interrupt)
	go func() {
		<-cc
		cancel()
		os.Exit(0)
	}()

	w.Add(1)
	go func() {
		for i := 0; i < c.Page["adPage"]; i++ {
			l, err := mcbbsad.FindPage(i, cookie, LimitGet)
			if err != nil {
				panic(err)
			}
			for _, v := range l {
				adl <- v
			}
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
			for _, v := range ad {
				adl <- v
			}
		}()
	}

	w.Wait()
	cancel()

	f, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	bs := bufio.NewReader(f)
	for {
		b, err := bs.ReadBytes('\n')
		if err != nil {
			if !errors.Is(err, io.EOF) {
				panic(err)
			}
		}
		if len(b) == 0 {
			break
		}
		var ad structs.McbbsAd
		err = json.Unmarshal(b, &ad)
		if err != nil {
			panic(err)
		}
		m[ad.Hash()] = ad
	}

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
		i := i
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
	m         map[string]structs.McbbsAd
)

type conifg struct {
	Page map[string]int
}

func init() {
	flag.IntVar(&threadInt, "thread", 8, "thread")
	flag.IntVar(&retry, "retry", 10, "retry")
	flag.IntVar(&sleepTime, "sleep", 3000, "sleep")
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

	m = make(map[string]structs.McbbsAd)
	bb, err := os.ReadFile("data.json")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return
		}
		panic(err)
	}
	err = json.Unmarshal(bb, &m)
	if err != nil {
		panic(err)
	}
}
