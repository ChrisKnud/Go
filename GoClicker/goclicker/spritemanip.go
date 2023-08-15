package goclicker

import "github.com/hajimehoshi/ebiten/v2"

// PUBLIC FUNCTIONS

/*
Centers a sprite object
scale is the value used in ebitens Translate.
*/
func CenterImage(sprite *Sprite, scaleX float64, scaleY float64) {
	// Get the sprite's width and height
	width := sprite.img.Bounds().Size().X
	height := sprite.img.Bounds().Size().Y

	// Set the x position of the sprites left side
	// and top y position
	sprite.posX[0] = WindowCenterX - float64(width)/2.0*scaleX
	sprite.posY[0] = WindowCenterY - float64(height)/2.0*scaleY
}

/*
Returns the values to in GeoM.Scale()
to achieve the wanted pixel size
*/
func ImageSizeToScale(width float64, heigth float64, wantedWidth float64, wantedHeight float64) (float64, float64) {
	// x = Scale value for x axis
	// y = Scale value for y axis
	x := wantedWidth / width
	y := wantedHeight / heigth

	return x, y
}

/*
Performs the sprite clicked animation
*/
func PressedAnimation(op *ebiten.DrawImageOptions) {
}
