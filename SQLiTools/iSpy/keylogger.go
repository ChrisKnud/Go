package iSpy

import (
	"github.com/eiannone/keyboard"
)

func Get_single_keystroke() rune {
	char, _, err := keyboard.GetSingleKey()
	if err != nil {
		panic(err)
	}

	return char
}
