package tengin

import (
	"github.com/gdamore/tcell/v3"
)

/**

tile provides character and style information
canvas dictates which tiles go where and layering order
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
	Char rune
	Fg   Color
	Bg   Color
}

func NewTile(char rune) Tile {
	return Tile{
		Char: char,
	}
}

func NewBlankTile() Tile {
	return Tile{
		Char: 0,
	}
}

func NewBgTile(bg Color) Tile {
	return Tile{
		Char: 0,
		Bg:   bg,
	}
}
