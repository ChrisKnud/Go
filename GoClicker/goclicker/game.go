package goclicker

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Update curson positions
	g.cursPosX, g.cursPosY = ebiten.CursorPosition()

	// Write your game's logical update.
	g.onceUpdate.Do(func() {
		InitGame(g)
	})

	RockieClicked(g)
	UpgradeClicked(g)

	// Goroutine that increments the score each second
	go func() {
		for {
			<-g.ticker.C
			ApplyMultiplier(g)
		}
	}()

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	g.onceDraw.Do(func() {
		InitBoard(g, screen)
	})
	text.Draw(screen, DynamicString("Score: ", g.score), scoreFont, 10, 30, color.White)
	text.Draw(screen, DynamicString("Multiplier: ", g.multiplier), scoreFont, 10, 30+int(ScoreFontsize), color.White)
	screen.DrawImage(g.sprite.img, &mainOp)

	ups := g.upgradables

	// Draw upgradables bar
	for i := 0; i < len(ups); i++ {
		DrawUpgrade(ups[i], screen, &ups[i].op)
		screen.DrawImage(ups[i].sprite.img, &ups[i].op)
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WindowWidth, WindowHeight
}
