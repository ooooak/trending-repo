package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"./db"
	"./types"
)

var allowedLangs = [...]string{
	"c%23", // c#
	"f%23", // f#
	"c", "go", "ocaml", "rust", "swift", "typescript",
	"c++", "erlang", "clojure", "haskell", "elm", "elixir",
	"common-lisp", "crystal", "css", "d", "dart", "html",
	"java", "julia", "kotlin", "lua", "nim", "php", "python",
	"ruby", "sass", "scala", "webassembly",
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func urls() []string {
	// const
	url := "https://github-trending-api.now.sh/repositories?language="
	var urls []string
	for _, lang := range allowedLangs {
		urls = append(urls, url+lang+"&since=daily")
	}
	return urls
}

func getURLContent(url string, ch chan<- types.Repos) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- nil
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- nil
		return
	}

	var data types.Repos
	err = json.Unmarshal(body, &data)
	check(err)

	ch <- data
}

func copyOutput() {
	path, err := filepath.Abs("./static/data.js")
	check(err)

	input, err := ioutil.ReadFile(db.FileName)
	check(err)

	output := []byte("var __output = " + string(input) + ";")
	err = ioutil.WriteFile(path, output, 0644)
	check(err)
}

func main() {
	records, err := db.Read()
	check(err)

	urls := urls()
	ch := make(chan types.Repos)
	for _, url := range urls {
		go getURLContent(url, ch)
	}

	for range urls {
		records.Consume(<-ch)
	}

	db.Write(records)

	copyOutput()
}
