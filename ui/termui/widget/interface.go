package widget

import (
	"github.com/iamquang95/stockterm/ui/termui/datacenter"
)

// Widget wrap a termui widget with UpdateData function
type Widget interface {
	GetWidget() interface{}
	UpdateData(datacenter.DataCenter)
}
