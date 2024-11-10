package webui

import "fmt"

type htmlButton struct {
	clickFunc func()
}

// Render displays the HTML button.
func (b htmlButton) Render() {
	fmt.Println("<button>I am a HTML button</button>")
}

// Click clicks the HTML button.
func (b htmlButton) Click() {
	b.clickFunc()
}

// OnClick changes the behavior of the button on being clicked.
func (b *htmlButton) OnClick(f func()) {
	b.clickFunc = f
}
