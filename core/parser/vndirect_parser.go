package parser

import (
	"encoding/json"
	"github.com/iamquang95/stockterm/schema"
	"sort"
	"time"
)

func ParseInDayStockData(date time.Time, resp []byte) ([]schema.PriceAtTime, error) {
	data := &dataStruct{}
	err := json.Unmarshal(resp, data)
	if err != nil {
		return nil, err
	}
	res := make([]schema.PriceAtTime, 0)
	hits := data.Data.Hits
	for _, kv := range hits {
		t, err := time.Parse("15:04:05", kv.Source.Time)
		if err != nil {
			return nil, err
		}
		timestamp := time.Date(date.Year(), date.Month(), date.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
		res = append(res, schema.PriceAtTime{
			Price: kv.Source.Last,
			Time:  timestamp,
		})
	}
	sort.Slice(res, func(l, r int) bool {
		return res[l].Time.Before(res[r].Time)
	})
	return res, nil
}

type dataStruct struct {
	Data struct {
		Hits []struct {
			Source struct {
				Last   float64 `json:"last"`
				Symbol string  `json:"symbol"`
				Time   string  `json:"time"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"data"`
}
