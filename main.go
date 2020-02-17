package main

import (
	"fmt"
	"github.com/iamquang95/stockterm/core/crawler"
)

func main() {
	//terminalui.Render()
	x, err := crawler.GetLastTradeDayStockDetail("ITA")
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
