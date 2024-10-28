package ui

type Button interface {
	Render()
	OnClick(func())
	Click()
}

type Dialog interface {
	Render()
	NewButton() Button
}
