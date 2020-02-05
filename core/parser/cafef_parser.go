package parser

import (
	"encoding/json"

	"github.com/iamquang95/stockterm/schema"
)

// ParseListStock parse http response from cafef to an array of Stock
// http://banggia.cafef.vn/stockhandler.ashx
func ParseListStock(resp string) ([]schema.Stock, error) {
	var stocks []schema.Stock
	err := json.Unmarshal([]byte(resp), &stocks)
	if err != nil {
		return nil, err
	}
	return stocks, nil
}
