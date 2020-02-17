package schema

import "time"

type Stock struct {
	Name       string  `json:"a"`
	OpenPrice  float32 `json:"b"`
	CeilPrice  float32 `json:"c"`
	FloorPrice float32 `json:"d"`
	Price      float32 `json:"l"`
}

type StockToday struct {
	Stock  Stock
	Prices []PriceAtTime
}

type PriceAtTime struct {
	Price float32
	Time time.Time
}
