package main

import (
	"strings"
	"fmt"
	"net/http"
	"regexp"
	"io/ioutil"
)

func searchUs(item string) {

	searchAmazon(item)
}

func searchAmazon(item string) {
	item = strings.Replace(item, " ", "+", -1)

	url := "https://www.amazon.com/s/?field-keywords=";

	url += item

	resp, _ := http.Get(url)

	defer resp.Body.Close()

	source, _ := ioutil.ReadAll(resp.Body)

	r := regexp.MustCompile(`<h2 data-attribute="(.*?)".*?href="(.*?)".*?aria-label="(\$.*?\...)`)
	results := r.FindAllSubmatch(source, -1)

	if results == nil {
		fmt.Println("nil")
		return;
	}

	for i := 0; i < len(results); i += 1 {
		fmt.Printf("%s\n", results[i][1])
		fmt.Printf("%s\n", results[i][2])
		fmt.Printf("%s\n", results[i][3])
	}
}