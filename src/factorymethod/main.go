package main

import (
	"fmt"

	"github.com/jcc620/DesignPatterns/src/factorymethod/ui"
	"github.com/jcc620/DesignPatterns/src/factorymethod/webui"
	"github.com/jcc620/DesignPatterns/src/factorymethod/windowsui"
)

func main() {
	webDialog := webui.WebDialog{}
	windowsDialog := windowsui.WindowsDialog{}

	playWithDialog(webDialog)
	fmt.Println("")
	playWithDialog(windowsDialog)
}

func playWithDialog(d ui.Dialog) {
	btn := d.NewButton()
	btn.Render()

	btn.Click()
	btn.OnClick(func() { fmt.Println("New Click Behaviour!") })
	btn.Click()
}
