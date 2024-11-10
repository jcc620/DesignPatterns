// Package webui provides implementation for web UI elements.
package webui

import (
	"fmt"

	"github.com/jcc620/DesignPatterns/src/factorymethod/ui"
)

// A HTMLButton is a clickable HTML button.
type HTMLButton struct {
	clickFunc func()
}

// Render displays the HTML button.
func (b HTMLButton) Render() {
	fmt.Println("I am a HTML button")
}

// Click clicks the HTML button.
func (b HTMLButton) Click() {
	b.clickFunc()
}

// OnClick changes the behavior of the button on being clicked.
func (b *HTMLButton) OnClick(f func()) {
	b.clickFunc = f
}

type WebDialog struct{}

func (d WebDialog) NewButton() ui.Button {
	defaultClickFunc := func() {
		fmt.Println("This HTML button was clicked!")
	}
	return &HTMLButton{clickFunc: defaultClickFunc}
}
