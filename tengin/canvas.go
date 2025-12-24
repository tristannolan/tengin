package tengin

import (
	"strings"
)

// Anything visual can be a canvas. Use it to draw images and text, set
// background colours, or group canvases together
type Canvas struct {
	X, Y, Z       int
	Width, Height int
	Tiles         [][]*Tile
	Children      []*Canvas
}

func NewCanvas(x, y, width, height int) *Canvas {
	tiles := make([][]*Tile, height)
	for i := range tiles {
		tiles[i] = make([]*Tile, width)
	}

	return &Canvas{
		X:        x,
		Y:        y,
		Z:        0,
		Width:    width,
		Height:   height,
		Tiles:    tiles,
		Children: []*Canvas{},
	}
}

func (c *Canvas) compose(offsetX, offsetY int, ops *[]*drawOp) {
	for y := range c.Tiles {
		for x := range c.Tiles[y] {
			opX := c.X + x + offsetX
			opY := c.Y + y + offsetY
			op := newDrawOp(opX, opY, c.Z, c.Tiles[y][x])
			*ops = append(*ops, &op)
		}
	}

	for i := range c.Children {
		c.Children[i].compose(c.X+offsetX, c.Y+offsetY, ops)
	}
}

func (c *Canvas) SetTile(x, y int, t *Tile) {
	if c.ContainsPoint(x, y) {
		c.Tiles[y][x] = t
	}
}

func (c Canvas) ContainsPoint(x, y int) bool {
	return x >= 0 &&
		x < c.Width &&
		y >= 0 &&
		y < c.Height
}

func (c *Canvas) AppendChild(children ...*Canvas) {
	for _, child := range children {
		child.Z += c.Z
		c.Children = append(c.Children, child)
	}
}

func (c *Canvas) Translate(x, y int) {
	c.X += x
	c.Y += y
}

func Box(x, y, width, height int, bg Color) *Canvas {
	c := NewCanvas(x, y, width, height)
	for y := range c.Tiles {
		for x := range c.Tiles[y] {
			tile := NewTile("", NewStyle().Bg(bg))
			c.SetTile(x, y, &tile)
		}
	}
	return c
}

func Text(x, y int, str string) *Canvas {
	c := NewCanvas(x, y, len(str), 1)
	i := 0
	for char := range strings.SplitSeq(str, "") {
		tile := NewTile(char, NewStyle())
		c.SetTile(i, 0, &tile)
		i++
	}
	return c
}

func Paragraph(x, y, width int, str string) *Canvas {
	var lines []string

	for p := range strings.SplitSeq(str, "\n") {
		// Preserve blank lines
		if len(p) == 0 {
			lines = append(lines, "")
			continue
		}

		lastIndex := 0

		for {
			if lastIndex+width >= len(p) {
				lines = append(lines,
					strings.TrimSpace(string(p[lastIndex:])),
				)
				break
			}

			i := lastIndex + width

			// Go back to last space
			for i > lastIndex && p[i] != ' ' {
				i--
			}

			// No space found, force a wrap
			if i == lastIndex {
				i += width
			}

			lines = append(lines,
				strings.TrimSpace(string(p[lastIndex:i])),
			)

			if i < len(p) && p[i] == ' ' {
				i++
			}
			lastIndex = i
		}
	}

	c := NewCanvas(x, y, width, len(lines))
	for i, line := range lines {
		chars := strings.Split(line, "")
		for j, char := range chars {
			tile := NewTile(char, NewStyle())
			c.SetTile(j, i, &tile)
		}
	}

	return c
}
