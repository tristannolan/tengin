package tengin

import (
	"github.com/gdamore/tcell/v3"
)

type Color struct {
	r, g, b int32
}

type Style struct {
	bg Color
	fg Color
}

type Tile struct {
	Char  string
	Style *Style
}

func NewColor(r, g, b int32) Color {
	return Color{
		r: r,
		g: g,
		b: b,
	}
}

func NewEmptyColor() Color {
	return Color{
		r: -1,
		g: -1,
		b: -1,
	}
}

func NewStyle() *Style {
	return &Style{
		bg: NewEmptyColor(),
		fg: NewEmptyColor(),
	}
}

func NewTile(char string, style *Style) *Tile {
	return &Tile{
		Char:  char,
		Style: style,
	}
}

func NewBlankTile() *Tile {
	return &Tile{
		Char:  "",
		Style: NewStyle(),
	}
}

func (c Color) IsEmpty() bool {
	return c.r == -1 || c.g == -1 || c.b == -1
}

func (c Color) IsEqualTo(clr Color) bool {
	return c.r == clr.r && c.g == clr.g && c.b == clr.b
}

func (c *Color) tcell() tcell.Color {
	if c.IsEmpty() {
		return tcell.ColorDefault
	}
	return tcell.NewRGBColor(c.r, c.g, c.b)
}

func (s *Style) Bg(c Color) Color {
	return s.bg
}

func (s *Style) Fg(c Color) Color {
	return s.fg
}

func (s *Style) NewBg(r, g, b int32) *Style {
	return s.SetBg(NewColor(r, g, b))
}

func (s *Style) NewFg(r, g, b int32) *Style {
	return s.SetFg(NewColor(r, g, b))
}

func (s Style) GetBg() Color {
	return s.bg
}

func (s Style) GetFg() Color {
	return s.fg
}

func (s *Style) SetBg(c Color) *Style {
	s.bg = c
	return s
}

func (s *Style) SetFg(c Color) *Style {
	s.fg = c
	return s
}

func (s *Style) CopyValues(from *Style) {
	s.bg = from.bg
	s.fg = from.fg
}
