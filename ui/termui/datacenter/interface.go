package datacenter

import (
	"github.com/iamquang95/stockterm/schema"
)

// DataCenter is a center place that store stock data
type DataCenter interface {
	GetStockList() []*schema.Stock
	GetStockLifeTimePrice(string) (*schema.StockLifeTimePrice, error)
	FetchData() error
	ModifyWatchingStock([]string) error
	GetWatchingStock() []string
}
