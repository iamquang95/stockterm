package widget

import (
	ui "github.com/gizak/termui/v3"
	"github.com/iamquang95/stockterm/ui/terminalui/datacenter"
)

// Widget wrap a terminalui widget with UpdateData function
type Widget interface {
	GetWidget() ui.Drawable
	UpdateData(datacenter.DataCenter) error
}
