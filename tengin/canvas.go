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
		Tiles:    tiles,
		children: []*Canvas{},
	}
}

func (c *Canvas) compose(ops *[]drawOp) {
	for y := range c.Tiles {
		for x := range c.Tiles[y] {
			op := NewDrawOp(c.X+x, c.Y+y, c.Z, c.Tiles[y][x])
			*ops = append(*ops, op)
		}
	}

	DebugLog("cy len", len(c.Tiles))
	DebugLog("cx len", len(c.Tiles[0]))

	// compose children as well
}

func (c *Canvas) SetTile(x, y int, t *Tile) {
	c.Tiles[y][x] = t
}
