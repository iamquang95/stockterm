package app

import (
	"github.com/iamquang95/stockterm/ui/termui/datacenter"
	"github.com/iamquang95/stockterm/ui/termui/widget"
)

type MainApp struct {
	widgets        []*widget.Widget
	dataCenter     *datacenter.DataCenter
	watchingStocks []string
}
