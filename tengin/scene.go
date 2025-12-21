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

func (s *Scene) render(screen tcell.Screen) {
	drawOps := []drawOp{}
	for _, c := range s.canvases {
		c.compose(&drawOps)
	}

	DebugLog("Len Canvases", len(s.canvases))

	slices.SortStableFunc(drawOps, func(a, b drawOp) int {
		return cmp.Compare(a.z, b.z)
	})

	screenWidth, screenHeight := screen.Size()
	clip := NewRect(0, 0, screenWidth-1, screenHeight-1)
	s.tiles = make([][]*Tile, screenHeight)
	for i := range s.tiles {
		s.tiles[i] = make([]*Tile, screenWidth)
	}

	for _, op := range drawOps {
		if !clip.Contains(op.x, op.y) {
			continue
		}

		s.tiles[op.y][op.x] = op.tile
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
