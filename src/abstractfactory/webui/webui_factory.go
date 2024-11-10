package webui

import (
	"fmt"

	"github.com/jcc620/DesignPatterns/src/abstractfactory/ui"
)

type WebFactory struct{}

func (f *WebFactory) NewButton() ui.Button {
	defaultClickFunc := func() {
		fmt.Println("This HTML button was clicked!")
	}
	return &htmlButton{clickFunc: defaultClickFunc}
}

func (f *WebFactory) NewCheckbox() ui.Checkbox {
	return &htmlCheckbox{checked: false}
}

func (f *WebFactory) NewHeading() ui.Heading {
	return &htmlHeading{
		text: "I am a HTML heading",
	}
}
