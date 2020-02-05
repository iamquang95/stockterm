package main

import (
	"fmt"

	"github.com/iamquang95/stockterm/core/crawler"
	"github.com/iamquang95/stockterm/core/parser"
)

func main() {
	url := "http://banggia.cafef.vn/stockhandler.ashx"
	html, err := crawler.GetHTML(url)
	if err != nil {
		panic(err)
	}
	stocks, err := parser.ParseListStock(html)
	if err != nil {
		panic(err)
	}
	for _, stock := range stocks {
		fmt.Println(stock)
	}
}
