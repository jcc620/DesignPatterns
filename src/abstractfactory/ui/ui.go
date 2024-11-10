// Package ui provides an interface for UI elements and the abstract UI factory.
package ui

// A Button is a clickable button element.
type Button interface {
	Render()
	OnClick(func())
	Click()
}

// A Checkbox is an element that can be toggled on and off.
type Checkbox interface {
	Render()
	Toggle()
	IsChecked() bool
}

// A Heading is an element to display heading text.
type Heading interface {
	Render()
	SetText(string)
}

// A Factory is a creator of UI elements.
type Factory interface {
	NewButton() Button
	NewCheckbox() Checkbox
	NewHeading() Heading
}
