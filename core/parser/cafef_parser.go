package parser

import (
	"encoding/json"

	"github.com/iamquang95/stockterm/schema"
)

// ParseListStock parses http response from cafef to an array of Stock
// http://banggia.cafef.vn/stockhandler.ashx
func ParseListStock(resp []byte) ([]schema.Stock, error) {
	var stocks []schema.Stock
	err := json.Unmarshal(resp, &stocks)
	if err != nil {
		return nil, err
	}
	return stocks, nil
}

type stockLifeTimePrice struct {
	Prices         [][]float32 `json:"prices"`
	RealTimePrices [][]float32 `json:"realtimePrice"`
}

// ParseStockLifeTimePrice parses http response from cafef showing life time price of a stock
// https://s.cafef.vn/ajax/StockChartV3.ashx?symbol=VRE
func ParseStockLifeTimePrice(resp []byte) (*schema.StockLifeTimePrice, error) {
	lifeTimePricesUnformated := &stockLifeTimePrice{}
	err := json.Unmarshal(resp, lifeTimePricesUnformated)
	if err != nil {
		return nil, err
	}
	prices := make([]schema.PriceAtEpoch, 0)
	realTimePrices := make([]schema.PriceAtEpoch, 0)
	for _, price := range lifeTimePricesUnformated.Prices {
		priceAtEpoch := schema.PriceAtEpoch{
			Epoch:  int(price[0]),
			Price:  price[1],
			Amount: price[2],
		}
		prices = append(prices, priceAtEpoch)
	}
	for _, price := range lifeTimePricesUnformated.RealTimePrices {
		priceAtEpoch := schema.PriceAtEpoch{
			Epoch:  int(price[0]),
			Price:  price[1],
			Amount: price[2],
		}
		realTimePrices = append(realTimePrices, priceAtEpoch)
	}
	lifeTimePrices := &schema.StockLifeTimePrice{
		Prices:         prices,
		RealTimePrices: realTimePrices,
	}
	return lifeTimePrices, nil
}
