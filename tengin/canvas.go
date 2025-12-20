package tengin

import (
	"github.com/gdamore/tcell/v3"
)

/**

tile provides character and style information
canvas dictates which tiles go where and layering
drawOperations flatten canvas data
scene compiles drawOperations into final image

*/

type Color struct {
	value tcell.Color
}

func NewColor(r, g, b int32) Color {
	return Color{
		value: tcell.NewRGBColor(r, g, b),
	}
}

// A blueprint to for colour and content
type Tile struct {
	Char string
	Fg   Color
	Bg   Color
}

func NewTile(char string) Tile {
	return Tile{
		Char: char,
	}
}

// Anything visual can be a canvas. Use it to draw images and text, set
// background colours, or group canvases together
type Canvas struct {
	X, Y, Z       int
	width, height int
	tiles         [][]*Tile
	children      []*Canvas
}

func NewCanvas(width, height int) Canvas {
	tiles := make([][]*Tile, height)
	for i := range tiles {
		tiles[i] = make([]*Tile, width)
	}

	return Canvas{
		X:        0,
		Y:        0,
		Z:        -1,
		width:    width,
		height:   height,
		tiles:    tiles,
		children: []*Canvas{},
	}
}

func (c *Canvas) LoopTiles(loop func(x, y int)) {
	for y := range c.tiles {
		for x := range c.tiles[y] {
			loop(x, y)
		}
	}
}

func (c *Canvas) SetTile(x, y int, t *Tile) {
	c.tiles[y][x] = t
}

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

func (s *Scene) render(screen tcell.Screen) {
	// Create an empty tile map
	w, h := screen.Size()
	s.tiles = make([][]*Tile, h)
	for i := range s.tiles {
		s.tiles[i] = make([]*Tile, w)
	}

	// Insert all canvas tiles into the tile map
	for _, c := range s.canvases {
		// Skip this canvas if it's off the screen
		if c.X > w || c.X+c.width < 0 ||
			c.Y > h || c.Y+c.height < 0 {
			continue
		}

		// Copy tiles over
		for y := range c.tiles {
			for x := range c.tiles[y] {
				// Crop if it's going off the screen
				if c.X+x < 0 || c.X+x >= w ||
					c.Y+y < 0 || c.Y+y >= h {
					continue
				}
				s.tiles[c.Y+y][c.X+x] = c.tiles[y][x]
			}
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
