package schema

// Stock describe a stock data, json based on cafef data
type Stock struct {
	Name       string  `json:"a"`
	OpenPrice  float32 `json:"b"`
	CeilPrice  float32 `json:"c"`
	FloorPrice float32 `json:"d"`
	Price      float32 `json:"i"`
}
