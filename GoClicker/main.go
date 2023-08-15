package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kriger452/GoClicker/goclicker"
)

func main() {
	game := &goclicker.Game{}
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(goclicker.WindowWidth, goclicker.WindowHeight)
	ebiten.SetWindowTitle(goclicker.WindowTitle)

	// TODO: Set window icon

	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
