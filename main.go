package main

import (
	"fmt"
	"github.com/iamquang95/stockterm/ui/terminalui/datacenter"
)

func main() {
	//terminalui.Render()
	dc, err := datacenter.NewStockDataCenter([]string{"ITA"})
	if err != nil {
		panic(err)
	}
	fmt.Println(dc)
	fmt.Println("----------------------------------")
	err = dc.FetchData()
	if err != nil {
		panic(err)
	}
	fmt.Println(dc.GetStockList())
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
