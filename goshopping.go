package main

import (
	"flag"
	"fmt"
	"os"
)

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
