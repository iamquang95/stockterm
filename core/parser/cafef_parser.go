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
	return stocks, err
}
