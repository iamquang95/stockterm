package widget

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/iamquang95/stockterm/schema"
	"github.com/iamquang95/stockterm/ui/terminalui/datacenter"
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
	plot.SetRect(0, 0, 50, 15)
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
	w.plot.Data = priceAtTimeToPlotData(stockDetail)
	return nil
}

func priceAtTimeToPlotData(stock *schema.StockToday) [][]float64 {
	n := 500
	start := 9*60*60
	end := 15*60*60
	step := (end-start)/n

	res := make([][]float64, 2)
	res[0] = make([]float64, n)
	res[1] = make([]float64, n)

	idx := 0
	prices := append(stock.Prices, schema.PriceAtTime{})
	copy(prices[1:], prices)
	prices[0] = schema.PriceAtTime{
		Price: stock.Stock.OpenPrice,
		Time: time.Time{},
	}
	prices = append(stock.Prices, schema.PriceAtTime{
		Price: 0,
		Time:  time.Now().AddDate(1995, 8, 15),
	})

	for i := 0; i <= n; i++ {
		res[0][i] = stock.Stock.OpenPrice
		curTime := start + step*i
		for idx < len(prices) && timeToInt(prices[idx].Time) <= curTime {
			idx++
		}
		res[1][i] = prices[idx-1].Price
	}

	return res
}

func timeToInt(t time.Time) int {
	return t.Hour()*60*60 + t.Minute()*60 + t.Second()
}
