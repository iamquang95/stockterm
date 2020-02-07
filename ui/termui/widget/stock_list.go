package widget

import (
	"fmt"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/iamquang95/stockterm/schema"
	"github.com/iamquang95/stockterm/ui/termui/datacenter"
)

type StockListWidget struct {
	table *widgets.Table
}

func NewStockListWidget() Widget {
	table := widgets.NewTable()
	table.TextStyle = ui.NewStyle(ui.ColorWhite)
	table.RowSeparator = true
	table.BorderStyle = ui.NewStyle(ui.ColorWhite)
	table.SetRect(0, 30, 70, 20)
	table.FillRow = true
	table.RowStyles[0] = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
	return &StockListWidget{
		table: table,
	}
}

func (w *StockListWidget) GetWidget() interface{} {
	return w.table
}

func (w *StockListWidget) UpdateData(dc datacenter.DataCenter) error {
	dataRows := [][]string{
		getTableHeader(),
	}
	stockList := dc.GetStockList()
	stockMap := make(map[string]*schema.Stock)
	for _, stock := range stockList {
		stockMap[stock.Name] = stock
	}
	for idx, code := range dc.GetWatchingStock() {
		stock, ok := stockMap[code]
		if !ok {
			return fmt.Errorf("Don't have data for %s", code)
		}
		entry := []string{
			stock.Name,
			floatToString(stock.Price),
		}
		dataRows = append(dataRows, entry)
		w.table.RowStyles[idx+1] = getRowStyle(stock)
	}
	w.table.Rows = dataRows
	return nil
}

func floatToString(x float32) string {
	return fmt.Sprintf("%.2f", x)
}

func getTableHeader() []string {
	return []string{"Code", "Price"}
}

func getRowStyle(stock *schema.Stock) ui.Style {
	var textColor ui.Color
	if stock.Price < stock.OpenPrice {
		textColor = ui.ColorRed
	} else if stock.Price > stock.OpenPrice {
		textColor = ui.ColorGreen
	} else {
		textColor = ui.ColorWhite
	}
	return ui.NewStyle(textColor)
}
