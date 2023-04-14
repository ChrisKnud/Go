package main

import (
	"fmt"
	"pstoolsmod/iSpy"
)

func main() {
	//fmt.Println(menu.Display_menu())
	//menu.Choose_action()

	char := iSpy.Get_single_keystroke()

	fmt.Printf("logged key: %v", char)
}
