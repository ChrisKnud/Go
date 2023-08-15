package goclicker

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

/*
Returns a string displaying
the current score
*/
func DynamicString(str string, score float64) string {
	return fmt.Sprintf("%s %0.2f", str, score)
}

/*
Checks if the cursor position
is inside of a sprite
*/
func CheckCursor(g *Game, sprite Sprite) bool {
	insideBoundaryX := false
	insideBoundaryY := false

	leftX := sprite.posX[0]
	rightX := sprite.posX[1]
	topY := sprite.posY[0]
	bottomY := sprite.posY[1]

	// Check if cursor is on the sprites x-axis
	if leftX <= float64(g.cursPosX) && float64(g.cursPosX) <= rightX {
		insideBoundaryX = true
	} else {
		insideBoundaryX = false
	}

	// Check if cursor is on the sprites y-axis
	if topY <= float64(g.cursPosY) && float64(g.cursPosY) <= bottomY {
		insideBoundaryY = true
	} else {
		insideBoundaryY = false
	}

	// Checks if the mouse cursor is inside the sprite
	if insideBoundaryX && insideBoundaryY {
		return true
	} else {
		return false
	}
}

func checkButtonPress(button ebiten.MouseButton) bool {
	if inpututil.IsMouseButtonJustPressed(button) {
		return true
	} else {
		return false
	}
}

/*
Rock main sprite clicked
Increments the game score
*/
func RockieClicked(g *Game) {
	if CheckCursor(g, g.sprite) && checkButtonPress(ebiten.MouseButton0) {
		g.score += g.scoreIncrement
		PressedAnimation(&mainOp)
	}
}

/*
Checks if the player has enought score to buy an upgrade.
If they have, the score incrementer will be increased.
If not they will need a higher score.
*/
func UpgradeClicked(g *Game) {
	up := g.upgradables

	for i := 0; i < len(up); i++ {
		if CheckCursor(g, up[i].sprite) && checkButtonPress(ebiten.MouseButton0) {
			if g.score >= up[i].cost {
				g.scoreIncrement += up[i].scoreIncrementer // Increase the score incrementer
				g.multiplier += up[i].multiplier           // Increase score multiplier
				g.score -= up[i].cost                      // Decrement the score by the upgrades cost
				costIncrease(&up[i])                       // Increases the cost of the upgrade
			} else {
				fmt.Printf("You need %f more score!", up[i].cost-g.score)
			}
		}
	}

}

/*
Increases cost of an upgrade once its bought
*/
func costIncrease(up *Upgradable) {
	up.cost *= 1.5
}

/*
Applies the scoremultiplier to the score
each second. Ebiten is set to 60 ticks per second
*/
func ApplyMultiplier(g *Game) {
	g.score += g.multiplier
}
