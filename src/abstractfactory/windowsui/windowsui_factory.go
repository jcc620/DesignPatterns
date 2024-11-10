package windowsui

import (
	"fmt"

	"github.com/jcc620/DesignPatterns/src/abstractfactory/ui"
)

type WindowsFactory struct{}

func (f *WindowsFactory) NewButton() ui.Button {
	defaultClickFunc := func() {
		fmt.Println("This Windows button was clicked!")
	}
	return &windowsButton{clickFunc: defaultClickFunc}
}

func (f *WindowsFactory) NewCheckbox() ui.Checkbox {
	return &windowsCheckbox{checked: false}
}

func (f *WindowsFactory) NewHeading() ui.Heading {
	return &windowsHeading{
		text: "I am a Windows heading",
	}
}
