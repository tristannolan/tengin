package tengin

import (
	"github.com/gdamore/tcell/v3"
)

// need a better solution for this
// styles should be handled nicer than this
// maybe a separate style struct? tcell's been right so far

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

func NewBlankTile() Tile {
	return Tile{
		Char: "",
	}
}

func NewBgTile(bg Color) Tile {
	return Tile{
		Char: "",
		Bg:   bg,
	}
}
