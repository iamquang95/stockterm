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
	prices, minVal, maxVal := priceAtTimeToPlotData(stockDetail)
	w.plot.Data = prices
	w.plot.MaxVal = maxVal - minVal
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
	prices = append(stock.Prices, schema.PriceAtTime{
		Price: 0,
		Time:  time.Now().AddDate( 0, 0, 15),
	})

	idx := 0

	maxVal := float64(0)
	minVal := 1e9
	for i := 0; i < n; i++ {
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
	// Refined data
	for i, _ := range res[0] {
		res[0][i] -= minVal
		res[1][i] = stock.Stock.OpenPrice - minVal
	}
	return res, minVal, maxVal
}

func timeToInt(t time.Time) int {
	return t.Hour()*60*60 + t.Minute()*60 + t.Second()
}
