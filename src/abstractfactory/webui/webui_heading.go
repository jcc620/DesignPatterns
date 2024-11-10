package webui

import "fmt"

type htmlHeading struct {
	text string
}

func (h *htmlHeading) Render() {
	fmt.Printf("<h1>%s</h1>\n", h.text)
}

func (h *htmlHeading) SetText(text string) {
	h.text = text
}
