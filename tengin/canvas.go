package tengin

/**

[
	[1,2,3],
	[4,5,6],
	[7,8,9],
]

*/

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
		Z:        -1,
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
