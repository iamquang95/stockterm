package crawler

import (
	"fmt"
	"github.com/iamquang95/stockterm/core/parser"
	"github.com/iamquang95/stockterm/schema"
	"time"
)

func getStockDetail(code string, date time.Time) ([]byte, error) {
	dateFormat := "2006-01-02"
	limit := 1000
	url := fmt.Sprintf(
		"https://finfo-api.vndirect.com.vn/v3/stocks/intraday/history"+
			"?symbols=%s"+
			"&sort=-time"+
			"&limit=%d"+
			"&fromDate=%s"+
			"&toDate=%s"+
			"&fields=symbol,last,time",
		code,
		limit,
		date.Format(dateFormat),
		date.Format(dateFormat),
	)
	data, err := GetHTML(url)
	return data, err
}

func GetLastTradeDayStockDetail(code string) ([]schema.PriceAtTime, error) {
	t := time.Now()
	if t.Weekday() == time.Sunday {
		t = t.AddDate(0, 0, -2)
	} else if t.Weekday() == time.Saturday {
		t = t.AddDate(0, 0, -1)
	} else if t.Hour() <= 9 {
		t = t.AddDate(0, 0, -1)
	}
	resp, err := getStockDetail(code, t)
	if err != nil {
		return nil, err
	}
	res, err := parser.ParseInDayStockData(t, resp)
	return res, err
}
