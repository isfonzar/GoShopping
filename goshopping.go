package main

import (
	"flag"
	"fmt"
)

func main() {
	item := flag.String("item", "", "Item to search")

	flag.Parse()

	search(*item)
}

func search(item string) {
	fmt.Println(item)
}

func searchAmazon() {

}