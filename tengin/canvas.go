package tengin

import (
	"strings"
)

type Canvas struct {
	X, Y, Z                int
	translateX, translateY int
	Width, Height          int
	Tiles                  [][]*Tile
	Children               []*Canvas
	parent                 *Canvas
	Clip                   bool
	dirty                  bool
	dirtyZ                 bool
	cachedDrawOps          []*drawOp
	alwaysCache            bool
	DebugName              string
}

func NewCanvas(x, y, width, height int) *Canvas {
	tiles := make([][]*Tile, height)
	for i := range tiles {
		tiles[i] = make([]*Tile, width)
	}

	return &Canvas{
		X:             x,
		Y:             y,
		Z:             0,
		translateX:    0,
		translateY:    0,
		Width:         width,
		Height:        height,
		Tiles:         tiles,
		Children:      []*Canvas{},
		parent:        nil,
		Clip:          false,
		dirty:         true,
		dirtyZ:        true,
		cachedDrawOps: []*drawOp{},
		alwaysCache:   false,
		DebugName:     "",
	}
}

func NewWrapperCanvas() *Canvas {
	return NewCanvas(0, 0, 0, 0)
}

func (c *Canvas) compose(ops *[]*drawOp) {
	c.composeClip(0, 0, ops, nil)
}

func (c *Canvas) composeClip(offsetX, offsetY int, ops *[]*drawOp, clip *Rect) {
	effX := c.X + c.translateX + offsetX
	effY := c.Y + c.translateY + offsetY
	effMaxX := effX + c.Width - 1
	effMaxY := effY + c.Height - 1

	if c.dirty == true || (c.alwaysCache && len(c.cachedDrawOps) > 0) {
		c.dirty = false
		c.cachedDrawOps = c.cachedDrawOps[:0]

		if c.Clip == true && clip == nil {
			newClip := NewRect(effX, effY, effMaxX, effMaxY)
			clip = &newClip
		} else if c.Clip == true {
			if effX > clip.minX {
				clip.minX = effX
			}
			if effY > clip.minY {
				clip.minY = effY
			}
			if effX+c.Width-1 < clip.maxX {
				clip.maxX = effMaxX
			}
			if effY+c.Height-1 < clip.maxY {
				clip.maxY = effMaxY
			}
		}

		for y := range c.Tiles {
			for x := range c.Tiles[y] {
				opX := effX + x
				opY := effY + y

				if clip == nil || clip.Contains(opX, opY) {
					op := newDrawOp(opX, opY, c.Z, c.Tiles[y][x])
					c.cachedDrawOps = append(c.cachedDrawOps, op)
				}
			}
		}
	}

	*ops = append(*ops, c.cachedDrawOps...)

	for i := range c.Children {
		c.Children[i].composeClip(effX, effY, ops, clip)
	}
}

func (c *Canvas) markDirty() {
	c.dirty = true
	if c.parent != nil {
		c.parent.markDirty()
	}
}

func (c *Canvas) markChildrenDirty() {
	for _, child := range c.Children {
		child.dirty = true
		child.markChildrenDirty()
	}
}

func (c *Canvas) IsDirty() bool {
	return c.dirty
}

func (c *Canvas) AlwaysCache() bool {
	return c.alwaysCache
}

func (c *Canvas) SetAlwaysCache(t bool) {
	c.alwaysCache = t
}

func (c *Canvas) SetTile(x, y int, t *Tile) {
	if !c.ContainsPoint(x, y) {
		return
	}

	c.markDirty()

	if c.Tiles[y][x] == nil {
		c.Tiles[y][x] = t
		return
	}

	if t.Style.bg.IsEmpty() && c.Tiles[y][x].Style != nil {
		t.Style.bg = c.Tiles[y][x].Style.bg
	}
	if t.Style.fg.IsEmpty() && c.Tiles[y][x].Style != nil {
		t.Style.fg = c.Tiles[y][x].Style.fg
	}

	c.Tiles[y][x] = t
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
		child.parent = c
		c.Children = append(c.Children, child)
	}
	c.markDirty()
}

func (c *Canvas) FlushChildren() {
	c.Children = c.Children[:0]
	c.markDirty()
}

func (c *Canvas) Position(x, y int) {
	if c.X == x && c.Y == y {
		return
	}
	c.X = x
	c.Y = y
	c.markDirty()
}

func (c *Canvas) Translate(x, y int) {
	c.translateX += x
	c.translateY += y
	c.markDirty()
	c.markChildrenDirty()
}

func (c *Canvas) GetTranslation() (int, int) {
	return c.translateX, c.translateY
}

func Box(x, y, width, height int, bg Color) *Canvas {
	c := NewCanvas(x, y, width, height)
	for y := range c.Tiles {
		for x := range c.Tiles[y] {
			tile := NewTile("", NewStyle().SetBg(bg))
			c.SetTile(x, y, tile)
		}
	}
	return c
}

func Text(x, y int, str string) *Canvas {
	c := NewCanvas(x, y, len(str), 1)
	i := 0
	for char := range strings.SplitSeq(str, "") {
		tile := NewTile(char, NewStyle())
		c.SetTile(i, 0, tile)
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
			c.SetTile(j, i, tile)
		}
	}

	return c
}
