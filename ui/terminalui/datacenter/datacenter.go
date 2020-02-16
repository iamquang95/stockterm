package datacenter

import (
	"fmt"
	"github.com/iamquang95/stockterm/core/crawler"
	"github.com/iamquang95/stockterm/core/parser"
	"github.com/iamquang95/stockterm/schema"
)

type StockDataCenter struct {
	stockList         []schema.StockToday
	watchingStocks    []string
}

func (dc *StockDataCenter) GetStockList() []schema.StockToday {
	return dc.stockList
}

func NewStockDataCenter(watchingStocks []string) DataCenter {
	dc := &StockDataCenter{
		stockList:         make([]schema.StockToday, 0),
		watchingStocks:    watchingStocks,
	}
	return dc
}

func (dc *StockDataCenter) FetchData() error {
	resp, err := crawler.GetHTML("http://banggia.cafef.vn/stockhandler.ashx")
	if err != nil {
		return err
	}
	stocks, err := parser.ParseListStock(resp)
	stockMap := make(map[string]schema.Stock, 0)
	for _, stock := range stocks {
		stockMap[stock.Name] = stock
	}
	if err != nil {
		return err
	}
	for idx, stock := range dc.stockList {
		newStock, ok := stockMap[stock.Stock.Name]
		if !ok {
			return fmt.Errorf("missing %s in cafef stock list", stock.Stock.Name)
		}
		stockPtr := &dc.stockList[idx]
		stockPtr.Stock = newStock
		//stockPtr.Prices = append(stockPtr.Prices, schema.PriceTime{
		//	Price: newStock.Price,
		//	Time:  time.Now(),
		//})
	}
	return nil
}

func (dc *StockDataCenter) ModifyWatchingStock(newWatching []string) error {
	dc.watchingStocks = newWatching
	err := dc.FetchData()
	return err
}

func (dc *StockDataCenter) GetWatchingStock() []string {
	return dc.watchingStocks
}

func (dc *StockDataCenter) initData() error {
	// TODO: implement me
	return nil
}
