package main

import (
	"github.com/iamquang95/stockterm/ui/terminalui"
)

func main() {
	err := terminalui.Render()
	if err != nil {
		panic(err)
	}
}
