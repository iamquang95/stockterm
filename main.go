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
	url = "https://s.cafef.vn/ajax/StockChartV3.ashx?symbol=VRE"
	html, err = crawler.GetHTML(url)
	if err != nil {
		panic(err)
	}
	lifeTimePrices, err := parser.ParseStockLifeTimePrice(html)
	if err != nil {
		panic(err)
	}
	fmt.Println(lifeTimePrices)
}
