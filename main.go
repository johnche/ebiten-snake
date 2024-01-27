package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(500, 500)
	ebiten.SetWindowTitle("Snake")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatalf("failed running game: %v", err)
	}
}
