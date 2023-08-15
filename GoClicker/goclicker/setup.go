package goclicker

import (
	"fmt"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// PUBLIC FUNCTIONS

/*
Initializes the game
*/
func InitGame(g *Game) {
	// Temporary variables
	var (
		mainSprite Sprite
		err        error
	)
	// Initialize tick
	g.ticker = time.NewTicker(1 * time.Second)

	// Initialize the upgradables array
	g.upgradables = []Upgradable{}

	// Get the main sprite image
	// and check for errors
	mainSprite.img = initImage("./assets/sprites/rockcookie.png")
	CheckForError(err)

	// Make window resizeable
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	// Initalize main sprite
	// and score increment
	g.sprite = mainSprite
	g.scoreIncrement = 1

	// Load font
	// and Initialize the main
	// sprite coordinates
	LoadFont(&scoreFont, ScoreFontsize)
	InitSpriteCoords(&g.sprite, 1, 1)
}

/*
Draws the game board
*/
func InitBoard(g *Game, screen *ebiten.Image) {
	sprite := g.sprite

	// Initialize rock cookie
	initDrawingOptions(&mainOp, 1, 1, true)
	CenterImage(&sprite, 1, 1)
	mainOp.GeoM.Translate(sprite.posX[0], sprite.posY[0])

	// Initialize upgrade bar
	initUpgradeBar(g)
}

/*
Sets the coordinates properties of the sprite
*/
func InitSpriteCoords(sprite *Sprite, scaleX float64, scaleY float64) {

	// Centers and inits posX[0] and posY[0]
	CenterImage(sprite, scaleX, scaleY)

	// Get sprite width and height
	sprite.width = sprite.img.Bounds().Size().X
	sprite.height = sprite.img.Bounds().Size().Y

	// Calculate the coordinates for the sprites
	// right x and bottom y
	sprite.posX[1] = sprite.posX[0] + float64(sprite.width)
	sprite.posY[1] = sprite.posY[0] + float64(sprite.height)
}

/*
Loads the desired textfont
*/
func LoadFont(txtFont *font.Face, size float64) {

	// Load a true type font
	// and check for error
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	CheckForError(err)

	// Set the font settings
	// and check for errors
	*txtFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    size,
		DPI:     FontDPI,
		Hinting: font.HintingVertical,
	})
	CheckForError(err)
}

// PRIVATE FUNCTIONS

/*
Initializes and image
*/
func initImage(imageUrl string) *ebiten.Image {
	// Get the main sprite image
	// and check for errors
	img, _, err := ebitenutil.NewImageFromFile(imageUrl)
	CheckForError(err)

	return img
}

/*
Initializes drawingoptions
Scale = true: Adjust image scale
Scale = false: Adjust image size
*/
func initDrawingOptions(op *ebiten.DrawImageOptions, width float64, height float64, scale bool) {
	op.GeoM.Reset() // Reset the DrawImageOptions.GeoM

	if scale {
		op.GeoM.Scale(width, height)
	} else if !scale {
		op.GeoM.Translate(width, height)
	}
}

/*
Initializes an upgradable
up = upgradable
*/
func initUpgradable(upImg *ebiten.Image, index int, upMultiplier float64, upScoreIncrementer float64, upCost float64, upName string, upDescription string) Upgradable {
	upOp := ebiten.DrawImageOptions{}

	// Sets the drawing options.
	// x = x start position
	// y = y start position
	//x := float64(TileSize * index)
	//y := Yoffset

	x := Yoffset
	y := float64(TileSize * index)

	// Upgradable width and height
	upWidth := upImg.Bounds().Size().X
	upHeight := upImg.Bounds().Size().Y

	// Get the scale values
	scaleX, scaleY := ImageSizeToScale(float64(upWidth), float64(upHeight), float64(TileSize), float64(TileSize))

	upOp.GeoM.Reset()
	upOp.GeoM.Scale(scaleX, scaleY)
	upOp.GeoM.Translate(x, y)

	upgradable := Upgradable{
		sprite: Sprite{
			img:    upImg,
			posX:   [2]float64{x, x + float64(TileSize)}, // Calculates the x position for the item
			posY:   [2]float64{y, y + float64(TileSize)},
			width:  upWidth,
			height: upHeight,
		},

		multiplier:       upMultiplier,
		scoreIncrementer: upScoreIncrementer,
		op:               upOp,
		cost:             upCost,
		name:             upName,
		description:      upDescription,
	}

	// Check if the upgradable is successfully made
	// and increment the upgradables counter
	if upgradable.name != "" {
		upCount++
	}
	fmt.Printf("x1: %f, x2: %f, y1: %f, y2: %f\n", upgradable.sprite.posX[0], upgradable.sprite.posX[1], upgradable.sprite.posY[0], upgradable.sprite.posY[1])
	return upgradable
}

/*
Initializes upgrade menu
*/
func initUpgradeBar(g *Game) {
	g.upgradables = []Upgradable{
		initUpgradable(initImage(upImageUrls["panther"]), 1, 0, 10, 10, "Panther", "Increases each tap by 10"),
		initUpgradable(initImage(upImageUrls["supermom"]), 2, 2, 0, 20, "Super mom", "Increases the score by 2 every second"),
		initUpgradable(initImage(upImageUrls["scooterkittens"]), 3, 0, 50, 10000, "Scooter kittens", "Increases each tap by 50"),
		initUpgradable(initImage(upImageUrls["chicken"]), 4, 10, 100, 15000, "Chicken", "Increases each tap by 100,\nIncreases multiplier by 10."),
	}
	println("upCount: ", upCount)
	if g.upgradables == nil {
		log.Fatal("Error: Upgradables list is nil.")
	}
}

/*
Pushes upgradable to the array inside the Game object
*/
func pushUpradable(g *Game, upgradable *Upgradable) {
	g.upgradables = append(g.upgradables, *upgradable)
}

/*
Not currently used
*/
func calcImageArea(sprite Sprite) (leftX float64, rightX float64, topY float64, bottomY float64) {
	spriteWidth := sprite.img.Bounds().Size().X
	spriteHeight := sprite.img.Bounds().Size().Y

	leftX = sprite.posX[0]
	topY = sprite.posY[0]
	rightX = leftX + float64(spriteWidth)
	bottomY = topY + float64(spriteHeight)

	return leftX, rightX, topY, bottomY
}
