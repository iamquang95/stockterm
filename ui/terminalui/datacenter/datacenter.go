package datacenter

import (
	"github.com/iamquang95/stockterm/schema"
)

type StockDataCenter struct {
	stockList         map[string]schema.StockToday
	watchingStocks    []string
}

func (dc *StockDataCenter) GetStockList() []schema.StockToday {
	res := make([]schema.StockToday, 0)
	for _, kv := range dc.stockList {
		res = append(res, kv)
	}
	return res
}

func NewStockDataCenter(watchingStocks []string) DataCenter {
	dc := &StockDataCenter{
		stockList:         make(map[string]schema.StockToday),
		watchingStocks:    watchingStocks,
	}
	return dc
}

func (dc *StockDataCenter) FetchData() error {
	//stockMap, err := crawler.GetCurrentStockInfo()
	//if err != nil {
	//	return err
	//}
	//for idx, stock := range dc.stockList {
	//	newStock, ok := stockMap[stock.Stock.Name]
	//	if !ok {
	//		return fmt.Errorf("missing %s in cafef stock list", stock.Stock.Name)
	//	}
	//	//stockPtr.Prices = append(stockPtr.Prices, schema.PriceTime{
	//	//	Price: newStock.Price,
	//	//	Time:  time.Now(),
	//	//})
	//}
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

	//for _, code := range dc.watchingStocks {
	//	data, err := crawler.GetLastTradeDayStockDetail(code)
	//	if err != nil {
	//		return err
	//	}
	//		return err
	//	}
	//	dc.stockList[code.]
	//}
	return nil
}
