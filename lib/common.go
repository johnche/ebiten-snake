package lib

type Coordinate struct {
	X int
	Y int
}

func (c Coordinate) Equal(c2 Coordinate) bool {
	return (c.X == c2.X) && (c.Y == c2.Y)
}

func (c Coordinate) Add(c2 Coordinate) Coordinate {
	return Coordinate{
		X: c.X + c2.X,
		Y: c.Y + c2.Y,
	}
}

type Area struct {
	Columns int
	Rows    int
}

func (c Coordinate) IsWithin(area Area) bool {
	return c.X >= 0 && c.X < area.Columns &&
		c.Y >= 0 && c.Y < area.Rows
}

func (c Coordinate) OverlapAny(coordinates []Coordinate) *Coordinate {
	for _, coordinate := range coordinates {
		if coordinate.Equal(c) {
			return &coordinate
		}
	}

	return nil
}
