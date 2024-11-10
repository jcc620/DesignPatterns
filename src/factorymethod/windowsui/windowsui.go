// Package windowsui provides implementation for Windows UI elements.
package windowsui

import (
	"fmt"

	"github.com/jcc620/DesignPatterns/src/factorymethod/ui"
)

// A HTMLButton is a clickable HTML button.
type windowsButton struct {
	clickFunc func()
}

// Render displays the HTML button.
func (b windowsButton) Render() {
	fmt.Println("I am a Windows button")
}

// Click clicks the HTML button.
func (b windowsButton) Click() {
	b.clickFunc()
}

// OnClick changes the behavior of the button on being clicked.
func (b *windowsButton) OnClick(f func()) {
	b.clickFunc = f
}

type WindowsDialog struct{}

func (d WindowsDialog) NewButton() ui.Button {
	defaultClickFunc := func() {
		fmt.Println("This Windows button was clicked!")
	}
	return &windowsButton{clickFunc: defaultClickFunc}
}
