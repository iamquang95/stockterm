package parser

import (
	"encoding/json"
	"fmt"
	"time"
)

func ParseInDayStockData(date time.Time, resp []byte) (map[time.Time]float32, error) {
	fmt.Println(string(resp))
	data := &dataStruct{}
	err := json.Unmarshal(resp, data)
	if err != nil {
		return nil, err
	}
	res := make(map[time.Time]float32)
	hits := data.Data.Hits
	for _, kv := range hits {
		t, err := time.Parse("15:04:05", kv.Source.Time)
		if err != nil {
			return nil, err
		}
		timestamp := time.Date(date.Year(), date.Month(), date.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
		res[timestamp] = kv.Source.Last
		fmt.Println(kv.Source)
	}
	return res, nil
}

type dataStruct struct {
	Data struct {
		Hits []struct {
			Source struct {
				Last   float32 `json:"last"`
				Symbol string  `json:"symbol"`
				Time   string  `json:"time"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"data"`
}
