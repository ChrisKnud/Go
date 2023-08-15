package goclicker

import (
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

const (
	// Main window constants
	WindowTitle   string  = "Rockie clicker"
	WindowWidth   int     = 1280
	WindowHeight  int     = 1280
	WindowCenterX float64 = float64(WindowWidth) / 2.0
	WindowCenterY float64 = float64(WindowHeight) / 2.0

	// Game constants
	TPS int = ebiten.DefaultTPS

	// Font constants
	ScoreFontsize     float64 = 36
	UpNameFont        float64 = 24
	UpDescriptionFont float64 = 12
	FontDPI           float64 = 72 // Dots per inch

	// Menu bar constants
	Yoffset     float64 = 32
	TileSize    int     = WindowWidth / 10
	TextPadding int     = 32
)

var (
	// Drawing options
	mainOp = ebiten.DrawImageOptions{}

	// Text variables
	scoreFont         font.Face
	upNameFont        font.Face
	updescriptionFont font.Face

	// Upgradables
	upCount     int // Number of upgradables
	upImageUrls = map[string]string{
		"panther":        "./assets/sprites/upgradables/up_panther.png",
		"supermom":       "./assets/sprites/upgradables/up_supermom.png",
		"scooterkittens": "./assets/sprites/upgradables/up_scooterkittens.png",
		"chicken":        "./assets/sprites/upgradables/up_chicken.png",
	}
)

// Game implements ebiten.Game interface.
type Game struct {
	sprite         Sprite
	upgradables    []Upgradable
	cursPosX       int
	cursPosY       int
	onceUpdate     sync.Once
	onceDraw       sync.Once
	score          float64
	scoreIncrement float64
	multiplier     float64
	ticker         *time.Ticker
}

// Sprite struct containing
// information about a sprite
type Sprite struct {
	/*
		posX[0] = left x coordinate
		posX[1] = right x coordinate
		posY[0] = top y coordinate
		posY[1] = bottom y coordinate
	*/
	img    *ebiten.Image
	posX   [2]float64
	posY   [2]float64
	width  int
	height int
}

type Upgradable struct {
	sprite           Sprite
	multiplier       float64
	scoreIncrementer float64 // Increments the scoreincrementer in Game(). Increases the score gotten per tap
	op               ebiten.DrawImageOptions
	cost             float64
	name             string
	description      string
}
