package termui

import (
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/iamquang95/stockterm/ui/termui/datacenter"
	"github.com/iamquang95/stockterm/ui/termui/widget"
)

type MainApp struct {
	widgets        []widget.Widget
	dataCenter     datacenter.DataCenter
	watchingStocks []string
	grid           *ui.Grid
}

func (app *MainApp) react() {
	defer ui.Close()

	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				app.grid.SetRect(0, 0, payload.Width, payload.Height)
				ui.Clear()
				ui.Render(app.grid)
			}
		case <-ticker:
			app.dataCenter.FetchData()
			for _, w := range app.widgets {
				w.UpdateData(app.dataCenter)
			}
			ui.Render(app.grid)
		}
	}
}

func Render() {
	app := initMainApp([]string{"VRE", "MSN", "ITA", "CTG"})
	app.react()
}

func initMainApp(watchingStocks []string) *MainApp {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	dc := datacenter.NewStockDataCenter(watchingStocks)
	stockList := widget.NewStockListWidget()
	widgets := []widget.Widget{stockList}
	for _, w := range widgets {
		w.UpdateData(dc)
	}
	// ui.Render(stockList.GetWidget())
	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)
	grid.Set(
		ui.NewRow(
			1.0/2, ui.NewCol(1.0/3, stockList.GetWidget()),
		),
	)
	app := &MainApp{
		widgets:        widgets,
		dataCenter:     dc,
		watchingStocks: watchingStocks,
		grid:           grid,
	}
	return app
}
