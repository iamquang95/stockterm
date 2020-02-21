package widget

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/iamquang95/stockterm/schema"
	"github.com/iamquang95/stockterm/ui/terminalui/datacenter"
	"math"
	"time"
)

type StockPriceChart struct {
	code string
	plot *widgets.Plot
}

func NewStockPriceChart(code string) Widget {
	plot := widgets.NewPlot()
	plot.Title = code
	plot.AxesColor = ui.ColorWhite
	plot.LineColors[0] = ui.ColorGreen
	plot.LineColors[1] = ui.ColorYellow
	plot.SetRect(0, 0, 80, 20)
	return &StockPriceChart{
		code: code,
		plot: plot,
	}
}

func (w *StockPriceChart) GetWidget() ui.Drawable {
	return w.plot
}

func (w *StockPriceChart) UpdateData(dc datacenter.DataCenter) error {
	stockDetail, err := dc.GetStockDetail(w.code)
	if err != nil {
		return err
	}
	prices, _, _ := priceAtTimeToPlotData(stockDetail)
	w.plot.Data = prices
	/* Don't know why color is applied to all charts
	if stockDetail.Stock.Price < stockDetail.Stock.OpenPrice {
		w.plot.LineColors[0] = ui.ColorRed
	} else if stockDetail.Stock.Price > stockDetail.Stock.OpenPrice {
		w.plot.LineColors[0] = ui.ColorGreen
	} else {
		w.plot.LineColors[0] = ui.ColorYellow
	}
	 */
	return nil
}

func priceAtTimeToPlotData(stock *schema.StockToday) ([][]float64, float64, float64) {
	n := 60
	start := 9 * 60 * 60
	end := 15 * 60 * 60
	step := (end - start) / n

	res := make([][]float64, 2)
	res[0] = make([]float64, n)
	res[1] = make([]float64, n)

	prices := append(stock.Prices, schema.PriceAtTime{})
	copy(prices[1:], prices)
	prices[0] = schema.PriceAtTime{
		Price: stock.Stock.OpenPrice,
		Time:  time.Time{},
	}
	now := time.Now()
	prices = append(stock.Prices, schema.PriceAtTime{
		Price: 0,
		Time:  time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 1e9-1, time.UTC),
	})

	idx := 0

	maxVal := float64(0)
	minVal := 1e9
	for i := 0; i < n; i++ {
		res[1][i] = stock.Stock.OpenPrice
		curTime := start + step*i
		for idx < len(prices) && timeToInt(prices[idx].Time) <= curTime {
			idx++
		}
		if idx == 0 {
			res[0][i] = prices[idx].Price
		} else {
			res[0][i] = prices[idx-1].Price
		}
		maxVal = math.Max(maxVal, res[0][i])
		minVal = math.Min(minVal, res[0][i])
	}
	return res, minVal, maxVal
}

func timeToInt(t time.Time) int {
	return t.Hour()*60*60 + t.Minute()*60 + t.Second()
}
