package datacenter

import (
	"fmt"

	"github.com/iamquang95/stockterm/core/crawler"
	"github.com/iamquang95/stockterm/core/parser"
	"github.com/iamquang95/stockterm/schema"
)

type StockDataCenter struct {
	stockList         []*schema.Stock
	lifeTimePricesMap map[string]*schema.StockLifeTimePrice
	watchingStocks    []string
}

func NewStockDataCenter(watchingStocks []string) DataCenter {
	dc := &StockDataCenter{
		stockList:         nil,
		lifeTimePricesMap: make(map[string]*schema.StockLifeTimePrice),
		watchingStocks:    watchingStocks,
	}
	dc.FetchData()
	return dc
}

func (dc *StockDataCenter) GetStockList() []*schema.Stock {
	return dc.stockList
}

func (dc *StockDataCenter) GetStockLifeTimePrice(code string) (*schema.StockLifeTimePrice, error) {
	val, ok := dc.lifeTimePricesMap[code]
	if !ok {
		return nil, fmt.Errorf("DataCenter doesn't has data for %s", code)
	}
	return val, nil
}

func (dc *StockDataCenter) FetchData() error {
	resp, err := crawler.GetHTML("http://banggia.cafef.vn/stockhandler.ashx")
	if err != nil {
		return err
	}
	stocks, err := parser.ParseListStock(resp)
	if err != nil {
		return err
	}
	dc.stockList = stocks
	// TODO: using channel to concurrent get these data
	// for _, stock := range dc.watchingStocks {
	// 	url := "https://s.cafef.vn/ajax/StockChartV3.ashx?symbol=" + stock
	// 	resp, err = crawler.GetHTML(url)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	lifeTimePrice, err := parser.ParseStockLifeTimePrice(resp)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	dc.lifeTimePricesMap[stock] = lifeTimePrice
	// }
	return nil
}

func (dc *StockDataCenter) ModifyWatchingStock(newWatchings []string) error {
	dc.watchingStocks = newWatchings
	err := dc.FetchData()
	return err
}

func (dc *StockDataCenter) GetWatchingStock() []string {
	return dc.watchingStocks
}
