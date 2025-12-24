package tengin

import (
	"github.com/gdamore/tcell/v3"
)

type Color struct {
	r, g, b int32
}

func NewEmptyColor() Color {
	return Color{
		r: -1,
		g: -1,
		b: -1,
	}
}

func (c Color) IsEmpty() bool {
	return c.r == -1 || c.g == -1 || c.b == -1
}

func NewColor(r, g, b int32) Color {
	return Color{
		r: r,
		g: g,
		b: b,
	}
}

func (c *Color) tcell() tcell.Color {
	if c.IsEmpty() {
		return tcell.ColorDefault
	}
	return tcell.NewRGBColor(c.r, c.g, c.b)
}

type Style struct {
	bg Color
	fg Color
}

func NewStyle() *Style {
	return &Style{
		bg: NewEmptyColor(),
		fg: NewEmptyColor(),
	}
}

func (s *Style) Bg(c Color) *Style {
	s.bg = c
	return s
}

func (s *Style) Fg(c Color) *Style {
	s.fg = c
	return s
}

func (s Style) GetBg() Color {
	return s.bg
}

func (s Style) GetFg() Color {
	return s.fg
}

// A blueprint to for colour and content
type Tile struct {
	Char  string
	Style *Style
}

func NewTile(char string, style *Style) Tile {
	return Tile{
		Char:  char,
		Style: style,
	}
}

func NewBlankTile() Tile {
	return Tile{
		Char:  "",
		Style: NewStyle(),
	}
}
