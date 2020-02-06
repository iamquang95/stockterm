package main

import (
	"fmt"
	"log"

	"github.com/iamquang95/stockterm/core/crawler"
	"github.com/iamquang95/stockterm/core/parser"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
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

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	p := widgets.NewParagraph()
	p.Text = "Hello World!"
	p.SetRect(0, 0, 25, 5)

	ui.Render(p)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}
