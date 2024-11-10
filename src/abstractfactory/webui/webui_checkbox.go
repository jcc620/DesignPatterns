package webui

import "fmt"

type htmlCheckbox struct {
	checked bool
}

func (c *htmlCheckbox) Render() {
	checkedStatus := ""
	if c.checked {
		checkedStatus = " checked"
	}
	fmt.Printf("<input type=\"checkbox\"%s/>\n", checkedStatus)
}

func (c *htmlCheckbox) Toggle() {
	c.checked = !(c.checked)
}

func (c *htmlCheckbox) IsChecked() bool {
	return c.checked
}
