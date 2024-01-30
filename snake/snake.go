package snake

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/johnche/ebiten-snake/lib"
)

var Step = map[ebiten.Key]lib.Coordinate{
	ebiten.KeyUp:    {X: 0, Y: -1},
	ebiten.KeyDown:  {X: 0, Y: 1},
	ebiten.KeyLeft:  {X: -1, Y: 0},
	ebiten.KeyRight: {X: 1, Y: 0},
}

type Snake struct {
	Head      Part
	HasEaten  bool
	direction ebiten.Key
	Length    int
	Color     color.Color

	// This is redundant to head, but an optimization for 1 less loop. worth?
	Positions []lib.Coordinate
}

func New(startingPoint lib.Coordinate, startingDirection ebiten.Key) *Snake {
	return &Snake{
		Head:      Part{Position: startingPoint},
		direction: startingDirection,
		Color:     color.RGBA{0, 0, 255, 255},
		HasEaten:  false,
		Length:    1,
		Positions: []lib.Coordinate{startingPoint},
	}
}

func (s *Snake) SetDirection(pressedKeys []ebiten.Key) bool {
	// Take first valid pressed key
	for _, pressedKey := range pressedKeys {
		if _, ok := Step[pressedKey]; ok {
			s.direction = pressedKey
			return true
		}
	}

	return false
}

func (s *Snake) Eat() {
	s.HasEaten = true
	s.Length += 1
}

func (s *Snake) Move() {
	s.Positions = s.Head.Move(
		s.Head.Position.Add(Step[s.direction]),
		s.HasEaten,
		[]lib.Coordinate{},
	)

	s.HasEaten = false
}

func (s *Snake) NextStep() lib.Coordinate {
	return s.Head.Position.Add(Step[s.direction])
}

// DEBUG only ===========================
func (s Snake) Direction() ebiten.Key {
	return s.direction
}
