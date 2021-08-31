package main

import (
	"encoding/json"
	"flag"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/xxmdhs/emeraldsledger/mcbbsad"
	"github.com/xxmdhs/emeraldsledger/structs"
	"github.com/xxmdhs/emeraldsledger/thread"
)

func main() {
	w := sync.WaitGroup{}
	adl := []structs.McbbsAd{}
	lock := sync.Mutex{}
	w.Add(1)
	go func() {
		for i := 0; i < c.Page["adPage"]; i++ {
			l, err := mcbbsad.FindPage(i, retry, cookie)
			if err != nil {
				panic(err)
			}
			lock.Lock()
			adl = append(adl, l...)
			lock.Unlock()
			time.Sleep(5000 * time.Millisecond)
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
			ad := threadFind(i, v)
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

	b, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("data.json", b, 0777)
	if err != nil {
		panic(err)
	}
}

func threadFind(tid, page int) []structs.McbbsAd {
	adl := []structs.McbbsAd{}
	l := sync.Mutex{}
	w := sync.WaitGroup{}

	a := 0
	for i := 0; i < page; i++ {
		w.Add(1)
		go func() {
			a++
			ad, err := thread.FindPage(tid, i, retry, 3000)
			if err != nil {
				panic(err)
			}
			l.Lock()
			adl = append(adl, ad...)
			l.Unlock()
			w.Done()

			time.Sleep(5000 * time.Millisecond)
		}()
		if a > threadInt {
			w.Wait()
			a = 0
			time.Sleep(5000 * time.Millisecond)
		}
	}
	w.Wait()

	return adl
}

var (
	threadInt int
	retry     int
	cookie    string
	c         conifg
)

type conifg struct {
	Page map[string]int
}

func init() {
	flag.IntVar(&threadInt, "thread", 3, "thread")
	flag.IntVar(&retry, "retry", 10, "retry")
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
