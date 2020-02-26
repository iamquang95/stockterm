package terminalui

import (
	"github.com/iamquang95/stockterm/core/config"
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/iamquang95/stockterm/ui/terminalui/datacenter"
	"github.com/iamquang95/stockterm/ui/terminalui/widget"
)

type MainApp struct {
	conf           *config.Config
	widgets        []widget.Widget
	dataCenter     datacenter.DataCenter
	watchingStocks []string
	grid           *ui.Grid
}

func (app *MainApp) react() {
	ui.Render(app.grid)
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
	conf, err := config.GetConfig("config.json")
	if err != nil {
		return err
	}
	app, err := initMainApp(conf)
	if err != nil {
		return err
	}
	app.react()
	return nil
}

func initMainApp(conf *config.Config) (*MainApp, error) {
	dc, err := datacenter.NewStockDataCenter(conf.WatchingStocks)
	if err != nil {
		return nil, err
	}
	renderingWidgets := make([]widget.Widget, 0)

	p, err := widget.NewStockPortfolio(dc, conf)
	if err != nil {
		return nil, err
	}
	renderingWidgets = append(renderingWidgets, p)

	// TODO: Watching stocks size may larger than 4
	for _, code := range conf.WatchingStocks {
		renderingWidgets = append(renderingWidgets, widget.NewStockPriceChart(code))
	}
	for _, w := range renderingWidgets {
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
			1.0,
			ui.NewCol(
				1.0/3,
				ui.NewCol(1.0, p.GetWidget()),
			),
			ui.NewCol(
				2.0/3,
				ui.NewRow(
					1.0/2,
					ui.NewCol(1.0/2, renderingWidgets[1].GetWidget()),
					ui.NewCol(1.0/2, renderingWidgets[2].GetWidget()),

				),
				ui.NewRow(
					1.0/2,
					ui.NewCol(1.0/2, renderingWidgets[3].GetWidget()),
					ui.NewCol(1.0/2, renderingWidgets[4].GetWidget()),

				),
			),
		),
	)
	app := &MainApp{
		conf:           conf,
		widgets:        renderingWidgets,
		dataCenter:     dc,
		watchingStocks: conf.WatchingStocks,
		grid:           grid,
	}
	return app, nil
}
