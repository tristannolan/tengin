package tengin

// Anything visual can be a canvas. Use it to draw images and text, set
// background colours, or group canvases together
type Canvas struct {
	X, Y, Z       int
	width, height int
	Tiles         [][]*Tile
	children      []*Canvas
}

func NewCanvas(x, y, width, height int) Canvas {
	tiles := make([][]*Tile, height)
	for i := range tiles {
		tiles[i] = make([]*Tile, width)
	}

	return Canvas{
		X:        x,
		Y:        y,
		Z:        0,
		width:    width,
		height:   height,
		Tiles:    tiles,
		children: []*Canvas{},
	}
}

func (c *Canvas) compose(offsetX, offsetY int, ops *[]*drawOp) {
	for y := range c.Tiles {
		for x := range c.Tiles[y] {
			opX := c.X + x + offsetX
			opY := c.Y + y + offsetY
			op := NewDrawOp(opX, opY, c.Z, c.Tiles[y][x])
			*ops = append(*ops, &op)
		}
	}

	for i := range c.children {
		c.children[i].compose(c.X+offsetX, c.Y+offsetY, ops)
	}
}

func (c *Canvas) SetTile(x, y int, t *Tile) {
	c.Tiles[y][x] = t
}

func (c *Canvas) AppendChild(child *Canvas) {
	child.Z += c.Z
	c.children = append(c.children, child)
}

func Box(x, y, width, height int, bg Color) Canvas {
	c := NewCanvas(x, y, width, height)
	for y := range c.Tiles {
		for x := range c.Tiles[y] {
			tile := NewBgTile(bg)
			c.SetTile(x, y, &tile)
		}
	}
	return c
}

func Text(x, y int, str string) Canvas {
	c := NewCanvas(x, y, len(str), 1)
	for i := range c.Tiles[0] {
		tile := NewTile(rune(str[i]))
		c.SetTile(i, 0, &tile)
	}
	return c
}
