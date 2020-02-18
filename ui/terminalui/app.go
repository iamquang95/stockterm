package terminalui

import (
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/iamquang95/stockterm/ui/terminalui/datacenter"
	"github.com/iamquang95/stockterm/ui/terminalui/widget"
)

type MainApp struct {
	widgets        []widget.Widget
	dataCenter     datacenter.DataCenter
	watchingStocks []string
	grid           *ui.Grid
}

func (app *MainApp) react() {
	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(10 * time.Second).C
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

func Render() error {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize terminalui: %v", err)
	}
	defer ui.Close()
	app, err := initMainApp([]string{"VRE", "MSN", "ITA", "CTG"})
	if err != nil {
		return err
	}
	app.react()
	return nil
}

func initMainApp(watchingStocks []string) (*MainApp, error) {

	dc, err := datacenter.NewStockDataCenter(watchingStocks)
	if err != nil {
		return nil, err
	}
	widgets := []widget.Widget{}
	for _, w := range widgets {
		err := w.UpdateData(dc)
		if err != nil {
			return nil, err
		}
	}
	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)
	grid.Set(
		ui.NewRow(
			1.0/2, ui.NewCol(1.0/3),
		),
	)
	app := &MainApp{
		widgets:        widgets,
		dataCenter:     dc,
		watchingStocks: watchingStocks,
		grid:           grid,
	}
	return app, nil
}
