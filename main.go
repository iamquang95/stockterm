package main

import (
	"fmt"
	"github.com/iamquang95/stockterm/core/crawler"
	"github.com/iamquang95/stockterm/core/parser"
	"time"
)

func main() {
	//terminalui.Render()
	date := time.Date(2020, time.February, 7, 1, 1, 1, 1, time.UTC)
	resp, err := crawler.GetStockDetail("ITA", date)
	if err != nil {
		panic(err)
	}
	x, err := parser.ParseInDayStockData(date, resp)
	if err != nil {
		panic(err)
	}
	fmt.Println(x)
	// resp, err := crawler.GetHTML("http://banggia.cafef.vn/stockhandler.ashx")
	// if err != nil {
	// 	panic(err)
	// }
	// stocks, err := parser.ParseListStock(resp)
	// fmt.Println("Err = ", err)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(stocks)
}
