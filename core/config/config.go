package config

type Config struct {
	WatchingStocks []string    `json:"watchingStocks"`
	Portfolio      []Portfolio `json:"portfolio"`
}

type Portfolio struct {
	Code     string  `json:"code"`
	NoStocks int     `json:"noStocks"`
	BuyPrice float64 `json:"buyPrice"`
}
