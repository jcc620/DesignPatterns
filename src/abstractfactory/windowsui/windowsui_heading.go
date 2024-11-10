package windowsui

import "fmt"

type windowsHeading struct {
	text string
}

func (h *windowsHeading) Render() {
	for _, char := range h.text {
		fmt.Printf("%c\u0332", char)
	}
	fmt.Print("\n")
}

func (h *windowsHeading) SetText(text string) {
	h.text = text
}
