package windowsui

import "fmt"

type windowsCheckbox struct {
	checked bool
}

func (c *windowsCheckbox) Render() {
	if c.checked {
		fmt.Println("\u2611")
	} else {
		fmt.Println("\u2610")
	}
}

func (c *windowsCheckbox) Toggle() {
	c.checked = !(c.checked)
}

func (c *windowsCheckbox) IsChecked() bool {
	return c.checked
}
