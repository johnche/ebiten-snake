package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/johnche/ebiten-snake/game"
)

func main() {
	ebiten.SetWindowSize(500, 500)
	ebiten.SetWindowTitle("Snake")
	if err := ebiten.RunGame(game.New()); err != nil {
		log.Fatalf("failed running game: %v", err)
	}
}
