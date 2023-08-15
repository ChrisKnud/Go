package goclicker

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func DrawSprite(img *ebiten.Image, screen *ebiten.Image, op ebiten.DrawImageOptions) {
	screen.DrawImage(img, &mainOp)
}

/*
Draws an upgradable object to the
upgradables menu
*/
func DrawUpgrade(up Upgradable, screen *ebiten.Image, op *ebiten.DrawImageOptions) {
	// Load the required fonts
	LoadFont(&upNameFont, UpNameFont)
	LoadFont(&updescriptionFont, UpDescriptionFont)

	costStr := fmt.Sprintf("Cost: %0.2f", up.cost)

	screen.DrawImage(up.sprite.img, op)
	text.Draw(screen, up.name, upNameFont, int(up.sprite.posX[1]), int(up.sprite.posY[0])+TextPadding, color.White)
	text.Draw(screen, up.description, updescriptionFont, int(up.sprite.posX[1]), int(up.sprite.posY[0])+TextPadding*2, color.White)
	text.Draw(screen, costStr, updescriptionFont, int(up.sprite.posX[1]), int(up.sprite.posY[0])+TextPadding*3, color.White)
}
