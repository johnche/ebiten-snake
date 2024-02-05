package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/johnche/ebiten-snake/game"
)

func main() {
	width := 1000
	height := 1000

	newGame := game.New(
		game.WithWidth(width),
		game.WithHeight(height),
		game.WithTPS(20),
	)

	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Snake")

	if err := ebiten.RunGame(newGame); err != nil {
		log.Fatalf("failed running game: %v", err)
	}
}
