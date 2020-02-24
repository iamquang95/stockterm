package widget

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/iamquang95/stockterm/core/config"
	"github.com/iamquang95/stockterm/schema"
	"github.com/iamquang95/stockterm/ui/terminalui/datacenter"
	"time"
)

type OneYearPortfolioChart struct {
	conf *config.Config
	plot *widgets.Plot
}

func NewOneYearPortfolioChart(conf *config.Config) Widget {
	plot := widgets.NewPlot()
	plot.Title = "Portfolio in 1Y"
	plot.AxesColor = ui.ColorWhite
	plot.LineColors[0] = ui.ColorGreen
	plot.LineColors[1] = ui.ColorYellow
	plot.SetRect(0, 0, 80, 20)
	return &OneYearPortfolioChart{
		conf: conf,
		plot: plot,
	}
}

func (w *OneYearPortfolioChart) GetWidget() ui.Drawable {
	return w.plot
}

func (w *OneYearPortfolioChart) UpdateData(dc datacenter.DataCenter) error {
	stocks := dc.GetOneYearStockList()
	prices, err := w.priceAtTimeToPlotData(stocks)
	if err != nil {
		return err
	}
	w.plot.Data = prices
	return nil
}

func (w *OneYearPortfolioChart) priceAtTimeToPlotData(stocks map[string][]schema.PriceAtTime) ([][]float64, error) {
	n := 60
	start := time.Now().AddDate(-1, 0, 0)
	end := time.Now()
	dayInterval := int(end.Sub(start).Hours()/24) / n

	res := make([][]float64, 2)
	res[0] = make([]float64, 0)
	res[1] = make([]float64, 0)

	totalAtBuyTime := 0.0
	for _, profile := range w.conf.Portfolio {
		totalAtBuyTime = totalAtBuyTime + float64(profile.NoStocks)*profile.BuyPrice
	}

	// The time complexity of this code is O(N*M*log(N)) where N is size of OneYearPrices, M is number of stock in portfolio
	// This code can be optimized to O(N*M) using a pointer to walk through the OneYearPrices list
	for t := start; t.Before(end); t = t.AddDate(0, 0, dayInterval) {
		dayValue := 0.0
		for _, profile := range w.conf.Portfolio {
			code := profile.Code
			price, err := find(t, stocks[code])
			if err != nil {
				return nil, err
			}
			dayValue = dayValue + price.Price*float64(profile.NoStocks)
		}
		res[0] = append(res[0], dayValue/1000)
		res[1] = append(res[1], totalAtBuyTime/1000)
	}

	return res, nil
}

func find(t time.Time, prices []schema.PriceAtTime) (*schema.PriceAtTime, error) {
	if len(prices) == 0 {
		return nil, fmt.Errorf("empty prices array")
	}
	l := 0
	r := len(prices) - 1
	res := 0
	for l <= r {
		m := l + (r-l)/2
		if prices[m].Time.Before(t) || prices[m].Time == t {
			res = m
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return &prices[res], nil
}
