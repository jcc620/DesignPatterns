// Package webui provides implementation for web UI elements.
package webui

import (
	"fmt"

	"github.com/jcc620/DesignPatterns/src/factorymethod/ui"
)

// A HTMLButton is a clickable HTML button.
type htmlButton struct {
	clickFunc func()
}

// Render displays the HTML button.
func (b htmlButton) Render() {
	fmt.Println("I am a HTML button")
}

// Click clicks the HTML button.
func (b htmlButton) Click() {
	b.clickFunc()
}

// OnClick changes the behavior of the button on being clicked.
func (b *htmlButton) OnClick(f func()) {
	b.clickFunc = f
}

type WebDialog struct{}

func (d WebDialog) NewButton() ui.Button {
	defaultClickFunc := func() {
		fmt.Println("This HTML button was clicked!")
	}
	return &htmlButton{clickFunc: defaultClickFunc}
}
