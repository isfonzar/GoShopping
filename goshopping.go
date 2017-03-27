package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"github.com/fatih/color"
)

type searchResult struct {
	name []byte
	link []byte
	price []byte
}

func main() {
	searchQuery := flag.String("search", "", "Search for a product")

	country := flag.String("country", "us", "Country to search")

	flag.Parse()

	if len(*searchQuery) == 0 {
		fmt.Println("Enter a item to search using: 'GoShopping -search item-to-search'")
		os.Exit(0);
	}

	search(*searchQuery, *country)
}

func search(item string, country string) {

	switch country {
		case "us":
			searchUs(item)
		default:
			fmt.Println("Country not found")
			os.Exit(0)
	}

}

func genericSearch(source []byte, pattern string) []searchResult {
	var resultArray []searchResult

	r := regexp.MustCompile(pattern)
	results := r.FindAllSubmatch(source, -1)

	if results == nil {
		return resultArray;
	}

	for i := 0; i < len(results); i += 1 {
		item := searchResult{name:results[i][1], price:results[i][3], link:results[i][2]}
		resultArray = append(resultArray, item)
	}

	return resultArray
}

func printList(results []searchResult) {
	for i := 0; i < len(results); i += 1 {
		color.Green("Name: %s\n", results[i].name)
		color.Red("Price: %s\n", results[i].price)
		color.White("Link: %s\n", results[i].link)
	}
}