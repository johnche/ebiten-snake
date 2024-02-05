package snake

import "github.com/johnche/ebiten-snake/lib"

type Part struct {
	Position lib.Coordinate
	Next     *Part
}

func (p *Part) Move(newPosition lib.Coordinate, grow bool, result []lib.Coordinate) []lib.Coordinate {
	oldPosition := p.Position

	p.Position = newPosition
	result = append(result, p.Position)

	if p.Next != nil {
		result = p.Next.Move(oldPosition, grow, result)
	} else if grow {
		p.Next = &Part{Position: oldPosition}
		result = append(result, p.Next.Position)
	}

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
