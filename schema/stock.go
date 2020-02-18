package schema

import "time"

type Stock struct {
	Name       string  `json:"a"`
	OpenPrice  float64 `json:"b"`
	CeilPrice  float64 `json:"c"`
	FloorPrice float64 `json:"d"`
	Price      float64 `json:"l"`
}

type StockToday struct {
	Stock  Stock
	Prices []PriceAtTime
}

type PriceAtTime struct {
	Price float64
	Time time.Time
}
