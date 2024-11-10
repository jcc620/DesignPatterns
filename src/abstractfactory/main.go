package main

import (
	"fmt"

	"github.com/jcc620/DesignPatterns/src/abstractfactory/ui"
	"github.com/jcc620/DesignPatterns/src/abstractfactory/webui"
	"github.com/jcc620/DesignPatterns/src/abstractfactory/windowsui"
)

func main() {
	useWindowsFactory := true

	var factory ui.Factory = &webui.WebFactory{}
	if useWindowsFactory {
		factory = &windowsui.WindowsFactory{}
	}
	playWithFactory(factory)
}

func playWithFactory(f ui.Factory) {
	btn := f.NewButton()
	btn.Render()
	btn.Click()
	btn.OnClick(func() { fmt.Println("New Click Behaviour!") })
	btn.Click()
	fmt.Println("")

	checkbox := f.NewCheckbox()
	checkbox.Render()
	checkbox.Toggle()
	checkbox.Render()
	fmt.Println("")

	heading := f.NewHeading()
	heading.Render()
	heading.SetText("Hello World")
	heading.Render()
}
