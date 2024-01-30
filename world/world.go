package world

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/johnche/ebiten-snake/lib"
	"github.com/johnche/ebiten-snake/snake"
)

// Make sure actors dont collide or overlap
type World struct {
	//Canvas     *ebiten.Image
	Boundaries         lib.Area
	LastAppleSpawnTime time.Time
	Apples             []lib.Coordinate
	Snake              *snake.Snake
}

func New(rows, cols int) *World {
	snake := snake.New(lib.Coordinate{
		//X: rand.Intn(cols),
		//Y: rand.Intn(rows),
		X: 0,
		Y: 5,
	}, ebiten.KeyRight)

	newWorld := &World{
		Boundaries:         lib.Area{Columns: cols, Rows: rows},
		Apples:             []lib.Coordinate{},
		LastAppleSpawnTime: time.Now(),
		Snake:              snake,
	}

	newWorld.NewApple()
	return newWorld
}

func (w *World) NewApple() {
	w.Apples = append(w.Apples, w.RandCoordinate())
	w.LastAppleSpawnTime = time.Now()
}

func (w *World) RemoveApple(targetApple lib.Coordinate) {
	var apples []lib.Coordinate

	for _, apple := range w.Apples {
		if !targetApple.Equal(apple) {
			apples = append(apples, apple)
		}
	}

	w.Apples = apples
}

func (w *World) RandCoordinate() lib.Coordinate {
	return lib.Coordinate{
		X: rand.Intn(w.Boundaries.Columns),
		Y: rand.Intn(w.Boundaries.Rows),
	}
}

func (w *World) Update(pressedKeys []ebiten.Key) error {
	w.Snake.SetDirection(pressedKeys)

	if !w.Snake.NextStep().IsWithin(w.Boundaries) {
		return fmt.Errorf("Snake hit the wall: %v", w.Snake.Head.Position)
	}

	if len(pressedKeys) > 0 {
		w.Snake.Move()
	}

	if apple := w.Snake.Head.Position.OverlapAny(w.Apples); apple != nil {
		w.Snake.Eat()
		w.RemoveApple(*apple)
		w.NewApple()
	}

	return nil
}
