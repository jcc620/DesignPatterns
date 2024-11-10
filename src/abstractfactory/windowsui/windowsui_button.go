package windowsui

import "fmt"

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
