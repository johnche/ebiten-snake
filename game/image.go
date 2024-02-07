package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/johnche/ebiten-snake/lib"
)

type DrawableBlock struct {
	position lib.Coordinate
	color    color.Color
}

type ImageCache struct {
	Width       int
	Height      int
	images      []*ebiten.Image
	colorBlocks []DrawableBlock
}

func NewImageCache(width, height int) *ImageCache {
	return &ImageCache{
		Width:  width,
		Height: height,
	}
}

func (i *ImageCache) RegisterActor(coordinates []lib.Coordinate, color color.Color) {
	for _, coordinate := range coordinates {
		i.colorBlocks = append(i.colorBlocks, DrawableBlock{position: coordinate, color: color})
	}
}

func (i *ImageCache) newImage() *ebiten.Image {
	return ebiten.NewImage(i.Width, i.Height)
}

func (i *ImageCache) updateCache() {
	if diff := len(i.colorBlocks) - len(i.images); diff > 0 {
		for range diff {
			i.images = append(i.images, i.newImage())
		}
	}
}

func (i *ImageCache) clearBlocks() {
	i.colorBlocks = []DrawableBlock{}
}

func (i *ImageCache) DrawImages(screen *ebiten.Image) {
	i.updateCache()
	for j, block := range i.colorBlocks {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(block.position.X*i.Width), float64(block.position.Y*i.Height))

		i.images[j].Clear()
		i.images[j].Fill(block.color)

		screen.DrawImage(i.images[j], op)
	}

	i.clearBlocks()
}
