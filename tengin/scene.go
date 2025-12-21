package tengin

import (
	"cmp"
	"slices"

	"github.com/gdamore/tcell/v3"
)

// A scene is used by the engine to render canvases. Only one scene should be
// provided to the renderer
type Scene struct {
	// x, y, z       int
	// width, height int
	canvases []*Canvas
	tiles    [][]*Tile
}

func newScene() Scene {
	return Scene{
		canvases: []*Canvas{},
		tiles:    [][]*Tile{},
	}
}

func (s *Scene) AppendCanvas(c *Canvas) {
	s.canvases = append(s.canvases, c)
}

// Draw operations are flattened canvases that scene will compose
type drawOp struct {
	x, y int
	z    int
	tile *Tile
}

func NewDrawOp(x, y, z int, tile *Tile) drawOp {
	return drawOp{
		x:    x,
		y:    y,
		z:    z,
		tile: tile,
	}
}

type layer struct {
	z       int
	drawOps []*drawOp
}

func newLayer(z int) layer {
	return layer{
		z:       z,
		drawOps: []*drawOp{},
	}
}

func (s *Scene) render(screen tcell.Screen) {
	layers := []*layer{}
	for _, c := range s.canvases {
		layer := newLayer(c.Z)
		c.compose(0, 0, &layer.drawOps)
		layers = append(layers, &layer)
	}

	for i := range layers {
		slices.SortStableFunc(layers[i].drawOps, func(a, b *drawOp) int {
			return cmp.Compare(a.z, b.z)
		})
	}

	slices.SortStableFunc(layers, func(a, b *layer) int {
		return cmp.Compare(a.z, b.z)
	})

	screenWidth, screenHeight := screen.Size()
	clip := NewRect(0, 0, screenWidth-1, screenHeight-1)
	s.tiles = make([][]*Tile, screenHeight)
	for i := range s.tiles {
		s.tiles[i] = make([]*Tile, screenWidth)
	}

	for i := range layers {
		for _, op := range layers[i].drawOps {
			if !clip.Contains(op.x, op.y) {
				continue
			}

			s.tiles[op.y][op.x] = op.tile
		}
	}

	for y, row := range s.tiles {
		for x, tile := range row {
			if tile == nil {
				continue
			}

			screen.Put(x, y, tile.Char, tcell.StyleDefault.Foreground(tile.Fg.value).Background(tile.Bg.value))
		}
	}
}
