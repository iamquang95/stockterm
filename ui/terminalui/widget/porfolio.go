package widget

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/iamquang95/stockterm/core/config"
	"github.com/iamquang95/stockterm/ui/terminalui/datacenter"
)

type StockPortfolio struct {
	stocks []config.Portfolio
	p      *widgets.Paragraph
	t      *widgets.Table
	oneYearPortfolioChart Widget
}

func NewStockPortfolio(dc datacenter.DataCenter, conf *config.Config) (*StockPortfolio, error) {
	p := widgets.NewParagraph()
	p.Title = "Portfolio"
	p.SetRect(0, 0, 60, 10)
	p.Text = newBalancePara(dc, conf.Portfolio)

	t := widgets.NewTable()
	t.Title = "Stock prices"
	t.SetRect(0, 0, 60, 10)

	portfolioChart := NewOneYearPortfolioChart(conf)
	err := portfolioChart.UpdateData(dc)
	if err != nil {
		return nil, err
	}

	stockPortfolio := &StockPortfolio{
		stocks: conf.Portfolio,
		p:      p,
		t:      t,
		oneYearPortfolioChart: portfolioChart,
	}
	stockPortfolio.updatePortfolioTable(dc, conf.Portfolio)
	return stockPortfolio, nil
}

func newBalancePara(dc datacenter.DataCenter, stocks []config.Portfolio) string {
	currentStockPrices := dc.GetStockList()
	totalBuy := 0.0
	balance := 0.0
	for _, stock := range stocks {
		totalBuy += float64(stock.NoStocks) * stock.BuyPrice * 1.0
		balance += float64(stock.NoStocks) * currentStockPrices[stock.Code].Stock.Price
	}
	balanceStr := fmt.Sprintf("%.2f", balance)
	return fmt.Sprintf(
		"Total Buy: %.2f\n"+
			"Current balance: %s\n",
		totalBuy,
		balanceStr,
	)
}

func (s *StockPortfolio) updatePortfolioTable(dc datacenter.DataCenter, stocks []config.Portfolio) {
	curPrices := dc.GetStockList()

	rows := make([][]string, 0)
	rows = append(rows, []string{
		"Code",
		"Quantity",
		"Buy price",
		"Cur price",
		"Profit",
	})
	rowColor := make(map[int]ui.Style)
	totalProfit := 0.0
	for idx, stock := range stocks {
		curStatus := curPrices[stock.Code]
		curPrice := curStatus.Stock.Price
		profit := (curPrice - stock.BuyPrice) * float64(stock.NoStocks)
		totalProfit += profit
		rows = append(rows, []string{
			stock.Code,
			fmt.Sprintf("%d", stock.NoStocks),
			fmt.Sprintf("%.2f", stock.BuyPrice),
			fmt.Sprintf("%.2f", curPrice),
			fmt.Sprintf("%.2f", profit),
		})
		rowColor[idx+1] = getColorBasedOnProfit(profit)
	}
	rows = append(rows, []string{
		"Total",
		"",
		"",
		"",
		fmt.Sprintf("%.2f", totalProfit),
	})
	rowColor[len(rows)-1] = getColorBasedOnProfit(totalProfit)
	s.t.Rows = rows
	s.t.RowStyles = rowColor
}

func getColorBasedOnProfit(profit float64) ui.Style {
	var color ui.Style
	if profit < 0 {
		color = ui.NewStyle(ui.ColorRed)
	} else if profit > 0 {
		color = ui.NewStyle(ui.ColorGreen)
	} else {
		color = ui.NewStyle(ui.ColorYellow)
	}
	return color
}

func (s *StockPortfolio) GetWidget() ui.Drawable {
	grid := ui.NewGrid()
	grid.Set(
		ui.NewRow(1.0/5, ui.NewCol(1, s.p)),
		ui.NewRow(1.0/2, ui.NewCol(1, s.t)),
		ui.NewRow(3.0/10, ui.NewCol(1, s.oneYearPortfolioChart.GetWidget())),
	)
	return grid
}

func (s *StockPortfolio) UpdateData(dc datacenter.DataCenter) error {
	s.p.Text = newBalancePara(dc, s.stocks)
	s.updatePortfolioTable(dc, s.stocks)
	err := s.oneYearPortfolioChart.UpdateData(dc)
	return err
}
