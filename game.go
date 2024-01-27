package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	ScreenX int
	ScreenY int

	Map   *Map
	Speed int
}

func NewGame() *Game {
	return &Game{
		ScreenX: 500,
		ScreenY: 500,
		Map:     NewMap(500, 500, 10, 10),
		Speed:   5,
	}
}

func (g *Game) Update() error {
	// set snake direction + spawn apples in time intervals
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.Map.Snake.SetDirection(ebiten.KeyUp)
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.Map.Snake.SetDirection(ebiten.KeyDown)
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.Map.Snake.SetDirection(ebiten.KeyLeft)
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.Map.Snake.SetDirection(ebiten.KeyRight)
	}

	// is snake next position legal? then snake.move()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Map.Draw()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.ScreenX, g.ScreenY
}
