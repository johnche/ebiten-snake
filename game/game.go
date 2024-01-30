package game

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/johnche/ebiten-snake/world"
)

type Game struct {
	width      int
	height     int
	rows       int
	cols       int
	imageCache *ImageCache

	world *world.World

	tps     int
	tick    time.Time
	running bool
}

type Option func(*Game)

func WithWidth(width int) Option {
	return func(g *Game) {
		g.width = width
	}
}

func WithHeight(height int) Option {
	return func(g *Game) {
		g.height = height
	}
}

func WithRows(rows int) Option {
	return func(g *Game) {
		g.rows = rows
	}
}

func WithCols(cols int) Option {
	return func(g *Game) {
		g.cols = cols
	}
}

func WithTPS(tps int) Option {
	return func(g *Game) {
		g.tps = tps
	}
}

func New(opts ...Option) *Game {
	game := &Game{
		width:  500,
		height: 500,
		rows:   50,
		cols:   50,
		tps:    5,

		tick:    time.Now(),
		running: true,
	}

	for _, opt := range opts {
		opt(game)
	}

	game.world = world.New(game.rows, game.cols)
	game.imageCache = NewImageCache(game.width/game.cols, game.height/game.rows)

	return game
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.width, g.height
}

func (g *Game) Update() error {
	if !g.running {
		// Nothing to do when game over
		return nil
	}

	pressedKeys := inpututil.AppendPressedKeys([]ebiten.Key{})
	if err := g.world.Update(pressedKeys); err != nil {
		fmt.Printf("%v\n", err)
		g.running = false
	}

	return nil
}

func (g *Game) GameOver(screen *ebiten.Image) {
	//ebitenutil.DebugPrintAt(screen, "Game over!", 0, 0)
	ebitenutil.DebugPrint(screen, "Game over!")
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.imageCache.RegisterActor(g.world.Snake.Positions, color.RGBA{0, 0, 255, 255})
	g.imageCache.RegisterActor(g.world.Apples, color.RGBA{255, 255, 255, 255})
	g.imageCache.DrawImages(screen)

	if !g.running {
		g.GameOver(screen)
		return
	}
}
