package datacenter

import (
	"github.com/iamquang95/stockterm/schema"
)

// DataCenter is a center place that store stock data
type DataCenter interface {
	GetStockList() map[string]*schema.StockToday
	GetStockDetail(code string) (*schema.StockToday, error)
	FetchData() error
	ModifyWatchingStock([]string) error
	GetWatchingStock() []string
}
