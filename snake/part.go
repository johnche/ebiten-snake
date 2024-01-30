package snake

import (
	"fmt"

	"github.com/johnche/ebiten-snake/lib"
)

type Part struct {
	Position lib.Coordinate
	Next     *Part
}

func (p *Part) Move(newPosition lib.Coordinate, grow bool, result []lib.Coordinate) []lib.Coordinate {
	oldPosition := p.Position

	p.Position = newPosition
	result = append(result, p.Position)

	if p.Next != nil {
		p.Move(oldPosition, grow, result)
	} else if grow {
		p.Next = &Part{Position: oldPosition}
		result = append(result, p.Next.Position)
		fmt.Printf("lastpart: %v\n", p)
	}

	fmt.Printf("result: %v\n", result)
	return result
}

func (p *Part) Crash(coordinate lib.Coordinate) bool {
	if p.Position.Equal(coordinate) {
		return true
	}

	if p.Next == nil {
		return false
	}

	return p.Next.Crash(coordinate)
}
