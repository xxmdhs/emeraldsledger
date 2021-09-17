package main

import (
	"encoding/json"
	"os"

	"github.com/xxmdhs/emeraldsledger/structs"
)

func main() {
	b, err := os.ReadFile("data.json")
	if err != nil {
		panic(err)
	}
	m := map[string]structs.McbbsAd{}
	err = json.Unmarshal(b, &m)
	if err != nil {
		panic(err)
	}
	newm := map[string]structs.McbbsAd{}
	for _, v := range m {
		newm[v.Hash()] = v
	}
	ff, err := os.Create("data.json")
	if err != nil {
		panic(err)
	}
	defer ff.Close()
	en := json.NewEncoder(ff)
	en.SetEscapeHTML(false)
	en.SetIndent("", "    ")
	err = en.Encode(newm)
	if err != nil {
		panic(err)
	}
}
