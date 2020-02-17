package crawler

import (
	"github.com/iamquang95/stockterm/core/parser"
	"github.com/iamquang95/stockterm/schema"
)

func GetCurrentStockInfo() (map[string]schema.Stock, error) {
	resp, err := GetHTML("http://banggia.cafef.vn/stockhandler.ashx")
	if err != nil {
		return nil, err
	}
	stocks, err := parser.ParseListStock(resp)
	stockMap := make(map[string]schema.Stock, 0)
	for _, stock := range stocks {
		stockMap[stock.Name] = stock
	}
	return stockMap, nil
}
