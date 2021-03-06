package datacenter

import (
	"fmt"
	"github.com/iamquang95/stockterm/core/crawler"
	"github.com/iamquang95/stockterm/schema"
	"time"
)

type StockDataCenter struct {
	stockList        map[string]*schema.StockToday
	oneYearStockList map[string][]schema.PriceAtTime
	watchingStocks   []string
}

func (dc *StockDataCenter) GetOneYearStockList() map[string][]schema.PriceAtTime {
	return dc.oneYearStockList
}

func (dc *StockDataCenter) GetStockList() map[string]*schema.StockToday {
	return dc.stockList
}

func (dc *StockDataCenter) GetStockDetail(code string) (*schema.StockToday, error) {
	res, ok := dc.stockList[code]
	if !ok {
		return nil, fmt.Errorf("data center don't have data for %s", code)
	}
	return res, nil
}

func NewStockDataCenter(watchingStocks []string) (DataCenter, error) {
	dc := &StockDataCenter{
		stockList:        make(map[string]*schema.StockToday),
		oneYearStockList: make(map[string][]schema.PriceAtTime),
		watchingStocks:   watchingStocks,
	}
	err := dc.initData()
	return dc, err
}

func (dc *StockDataCenter) FetchData() error {
	stockMap, err := crawler.GetCurrentStockInfo()
	if err != nil {
		return err
	}
	for _, stock := range dc.watchingStocks {
		newStock, ok := stockMap[stock]
		if !ok {
			return fmt.Errorf("missing %s in cafef stock list", stock)
		}
		stockDetail, ok := dc.stockList[stock]
		if !ok {
			return fmt.Errorf("missing %s in data center stock list", stock)
		}
		if len(stockDetail.Prices) == 0 {
			continue
		}
		curYear, curMonth, curDay := time.Now().Date()
		lastYear, lastMonth, lastDay := stockDetail.Prices[0].Time.Date()
		if curYear == lastYear && curMonth == lastMonth && curDay == lastDay {
			stockDetail.Stock = newStock
			stockDetail.Prices = append(stockDetail.Prices, schema.PriceAtTime{
				Price: newStock.Price,
				Time:  time.Now(),
			})
		}
		oneYearStat := dc.oneYearStockList[stock]
		dc.oneYearStockList[stock][len(oneYearStat)-1] = schema.PriceAtTime{
			Price: newStock.Price,
			Time:  time.Now(),
		}
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
	stocks, err := crawler.GetCurrentStockInfo()
	if err != nil {
		return nil
	}
	for _, code := range dc.watchingStocks {
		data, err := crawler.GetLastTradeDayStockDetail(code)
		if err != nil {
			return err
		}
		dc.stockList[code] = &schema.StockToday{
			Stock:  stocks[code],
			Prices: data,
		}
		data, err = crawler.GetOneYearStockPrice(code)
		dc.oneYearStockList[code] = data
	}
	return nil
}
