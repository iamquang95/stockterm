package parser

import (
	"encoding/json"
	"github.com/iamquang95/stockterm/schema"
	"sort"
	"time"
)

func ParseInDayStockData(date time.Time, resp []byte) ([]schema.PriceAtTime, error) {
	data := &inDayDataStruct{}
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

type inDayDataStruct struct {
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

func ParseOneYearStockData(resp []byte) ([]schema.PriceAtTime, error) {
	data := &oneYearDataStruct{}
	err := json.Unmarshal(resp, data)
	if err != nil {
		return nil, err
	}
	res := make([]schema.PriceAtTime, 0)
	for i := 0; i < min(len(data.Epochs), len(data.Prices)); i = i + 1 {
		res = append(res, schema.PriceAtTime{
			Price: data.Prices[i],
			Time:  time.Unix(data.Epochs[i], 0),
		})
	}
	sort.Slice(res, func(l, r int) bool {
		return res[l].Time.Before(res[r].Time)
	})
	return res, nil
}

type oneYearDataStruct struct {
	Epochs []int64   `json:"t"`
	Prices []float64 `json:"c"`
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
