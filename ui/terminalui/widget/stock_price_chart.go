package widget

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/iamquang95/stockterm/ui/terminalui/datacenter"
)

type StockPriceChart struct {
	plot *widgets.Plot
}

func NewStockPriceChart() Widget {
	plot := widgets.NewPlot()
	plot.AxesColor = ui.ColorWhite
	plot.LineColors[0] = ui.ColorGreen
	plot.LineColors[1] = ui.ColorYellow
	return &StockPriceChart{
		plot: plot,
	}
}

func (w *StockPriceChart) GetWidget() ui.Drawable {
	return w.plot
}

func (w *StockPriceChart) UpdateData(dc datacenter.DataCenter) error {
	return nil
}
