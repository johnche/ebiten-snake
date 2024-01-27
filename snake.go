package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type SnakePart struct {
	Position Coordinate
}

type Snake struct {
	Parts     []SnakePart
	direction ebiten.Key
	Color     color.Color
}

const DirectionalStep = map[ebiten.Key]Coordinate{
	ebiten.KeyUp: Coordinate{X: 0, Y: -1},
	ebiten.KeyDown: Coordinate{X: 0, Y: 1},
	ebiten.KeyLeft: Coordinate{X: -1, Y: 0},
	ebiten.KeyRight: Coordinate{X: 1, Y: 0},
}

func NewSnake(startingPoint Coordinate, startingDirection ebiten.Key) *Snake {
	return &Snake{
		Parts:     []SnakePart{SnakePart{startingPoint}},
		direction: startingDirection,
		Color:     color.RGBA{0, 0, 255, 255},
	}
}

func (s Snake) SetDirection(direction ebiten.Key) {
	s.direction = direction
}

func (s *Snake) Eat() {
	// new snakepart should have coordinate
	s.Parts = append(s.Parts, SnakePart{})
}

func (s Snake) Move() {
	let previousCoordinate Coordinate
	for _, part := range s.Parts {
		if previousCoordinate == nil {
			previousCoordinate = part.Coordinate

			part.Coordinate = part.Coordinate.Add(DirectionalStep[s.direction])
		}
	}
}

func (s *Snake) CrashWith(p Coordinate) bool {
	for _, part := range s.Parts {
		if part.Position.Equal(p) {
			return true
		}
	}

	return false
}
