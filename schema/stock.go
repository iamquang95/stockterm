package schema

// Stock describes a stock data, json based on cafef data
type Stock struct {
	Name       string  `json:"a"`
	OpenPrice  float32 `json:"b"`
	CeilPrice  float32 `json:"c"`
	FloorPrice float32 `json:"d"`
	Price      float32 `json:"i"`
}

// StockLifeTimePrice decribes a life time price of a stock
// Prices is price at end of the day of this stock per day
// RealTimePrices shows the price adjustment of this stock on the last day
type StockLifeTimePrice struct {
	Prices         []PriceAtEpoch
	RealTimePrices []PriceAtEpoch
}

// PriceAtEpoch describes a stock price at an epoch with total amount
type PriceAtEpoch struct {
	Epoch  int
	Price  float32
	Amount float32
}
