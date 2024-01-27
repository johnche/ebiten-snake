package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// Make sure actors dont collide or overlap
type Map struct {
	Canvas  *ebiten.Image
	Columns int
	Rows    int

	TileSizeX int
	TileSizeY int

	LastAppleSpawnTime time.Time
	Snake              *Snake

	Apples []Apple
}

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

type Apple struct {
	Position Coordinate
}

func (m *Map) newApple() {
	for true {
		newCoordinate := Coordinate{
			X: rand.Intn(m.Columns),
			Y: rand.Intn(m.Rows),
		}

		if !m.IsIllegalPosition(newCoordinate) {
			m.Apples = append(m.Apples, Apple{newCoordinate})
			m.LastAppleSpawnTime = time.Now()
			return
		}
	}
}

func NewMap(sizeX, sizeY, cols, rows int) *Map {
	snake := NewSnake(Coordinate{
		X: rand.Intn(cols),
		Y: rand.Intn(rows),
	}, ebiten.KeyRight)

	newMap := &Map{
		Canvas:  ebiten.NewImage(sizeX, sizeY),
		Columns: cols,
		Rows:    rows,

		TileSizeX: sizeX / cols,
		TileSizeY: sizeY / rows,

		Apples:             []Apple{},
		LastAppleSpawnTime: time.Now(),

		Snake: snake,
	}

	newMap.newApple()
	return newMap
}

func (m *Map) CreateRandomCoordinate() Coordinate {
	for true {
		newCoordinate := Coordinate{
			X: rand.Intn(m.Columns),
			Y: rand.Intn(m.Rows),
		}

		if !m.IsIllegalPosition(newCoordinate) {
			return newCoordinate
		}
	}
}

func (m *Map) IsIllegalPosition(position Coordinate) bool {
	for _, apple := range m.Apples {
		if apple.Position.Equal(position) {
			return true
		}
	}

	return m.Snake.CrashWith(position)
}

func (m *Map) Draw() {
	for _, apple := range m.Apples {
		m.RenderItem(apple.Position, color.RGBA{0, 0, 0, 255})
	}

	for _, snakePart := range m.Snake.Parts {
		m.RenderItem(snakePart.Position, m.Snake.Color)
	}
}

func (m *Map) RenderItem(p Coordinate, c color.Color) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.X*m.TileSizeX), float64(p.Y*m.TileSizeY))

	image := ebiten.NewImage(m.TileSizeX, m.TileSizeY)
	image.Fill(c)

	m.Canvas.DrawImage(image, op)
}
