// Package ui provides an interface for UI elements.
package ui

// A Button is a clickable button element.
type Button interface {
	Render()
	OnClick(func())
	Click()
}

// A Dialog is a container for UI elements.
type Dialog interface {
	NewButton() Button
}
