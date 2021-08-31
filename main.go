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
	if gen {
		save()
		return
	}

	w := sync.WaitGroup{}
	adl := make(chan structs.McbbsAd, 50)
	cxt, cancel := context.WithCancel(context.Background())

	LimitGet := http.NewLimitGet(threadInt, sleepTime, retry)

	go func() {
		defer w.Done()
		f, err := os.Create("data.txt")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		for {
			select {
			case <-cxt.Done():
				return
			case ad := <-adl:
				b, err := json.Marshal(ad)
				if err != nil {
					panic(err)
				}
				f.Write(b)
				f.Write([]byte("\n"))
				log.Println(string(b))
				w.Done()
			}
		}
	}()

	w.Add(1)
	go func() {
		for i := 0; i < c.Page["adPage"]; i++ {
			l, err := mcbbsad.FindPage(i, cookie, LimitGet)
			if err != nil {
				panic(err)
			}
			for _, v := range l {
				w.Add(1)
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
			threadFind(i, v, LimitGet, adl, &w)
			w.Done()
		}()
	}

	cc := make(chan os.Signal, 1)
	signal.Notify(cc, os.Interrupt)
	go func() {
		<-cc
		w.Add(1)
		cancel()
		w.Wait()
		save()
		os.Exit(0)
	}()

	w.Wait()

	w.Add(1)
	cancel()
	w.Wait()

	save()
}

func save() {
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
	ff, err := os.Create("data.json")
	if err != nil {
		panic(err)
	}
	defer ff.Close()
	en := json.NewEncoder(ff)
	en.SetEscapeHTML(false)
	en.SetIndent("", "    ")
	err = en.Encode(m)
	if err != nil {
		panic(err)
	}
}

func threadFind(tid, page int, LimitGet *http.LimitGet, ch chan<- structs.McbbsAd, w *sync.WaitGroup) {
	ww := sync.WaitGroup{}

	a := 0
	for i := 0; i < page; i++ {
		ww.Add(1)
		i := i
		go func() {
			a++
			ad, err := thread.FindPage(tid, i, LimitGet)
			if err != nil {
				panic(err)
			}
			for _, v := range ad {
				w.Add(1)
				ch <- v
			}
			ww.Done()
		}()
		if a > threadInt {
			ww.Wait()
			a = 0
		}
	}
	ww.Wait()
}

var (
	threadInt int
	retry     int
	cookie    string
	sleepTime int
	c         conifg
	m         map[string]structs.McbbsAd
	gen       bool
)

type conifg struct {
	Page map[string]int
}

func init() {
	flag.IntVar(&threadInt, "thread", 8, "thread")
	flag.IntVar(&retry, "retry", 10, "retry")
	flag.IntVar(&sleepTime, "sleep", 3000, "sleep")
	flag.BoolVar(&gen, "gen", false, "gen")
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
