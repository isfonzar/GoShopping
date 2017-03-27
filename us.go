package main

import (
	"strings"
	"net/http"
	"io/ioutil"
)

func searchUs(item string) {
	results := searchAmazon(item)

	printList(results)
}

func searchAmazon(item string) []searchResult {
	item = strings.Replace(item, " ", "+", -1)

	url := "https://www.amazon.com/s/?field-keywords=";

	url += item

	resp, _ := http.Get(url)

	defer resp.Body.Close()

	source, _ := ioutil.ReadAll(resp.Body)

	return genericSearch(source, `<h2 data-attribute="(.*?)".*?href="(http.*?)".*?aria-label="(\$.*?\...)`)
}