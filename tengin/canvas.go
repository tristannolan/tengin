package tengin

import "errors"

type Canvas struct {
	control       ControlManager
	x, y, z       int
	transform     *Transform
	Width, Height int
	Tiles         [][]*Tile
	Children      []*Canvas
	parent        *Canvas
	Clip          bool
	dirty         bool
	dirtyZ        bool
	cachedDrawOps []*drawOp
	alwaysCache   bool
	DebugName     string
}

var ErrCanvasDoesNotContainPoint = errors.New("Canvas - Does not contain point")

// Don't render tiles if you don't have to. An empty canvas can display child
// tiles if clipping is disabled.
func NewWrapperCanvas() *Canvas {
	return NewCanvas(0, 0)
}

// The general purpose rendering medium. A canvas is a locally scoped tile map.
// It renders tiles provided, and can nest canvases to create complex images.
// A canvas does not inherently have a style, it's literally a blank canvas.
// Style information can be set directly onto a tile.
// Use the transform property to position a canvas in the world view.
// The local position will be set relative to the parent canvas, and ultimately
// the scene that it is rendered to.
func NewCanvas(width, height int) *Canvas {
	tiles := make([][]*Tile, height)
	for i := range tiles {
		tiles[i] = make([]*Tile, width)
	}

	return &Canvas{
		x:             0,
		y:             0,
		z:             0,
		transform:     NewTransform(0, 0),
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

func (c *Canvas) compose(ops *[]*drawOp) {
	c.composeClip(0, 0, ops, nil)
}

func (c *Canvas) composeClip(offsetX, offsetY int, ops *[]*drawOp, clip *Rect) {
	effX := c.x + c.transform.x + offsetX
	effY := c.y + c.transform.y + offsetY
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
					op := newDrawOp(opX, opY, c.z, c.Tiles[y][x])
					c.cachedDrawOps = append(c.cachedDrawOps, op)
				}
			}
		}
		*ops = append(*ops, c.cachedDrawOps...)
	}

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

func (c *Canvas) Tile(x, y int) (*Tile, error) {
	if !c.ContainsPoint(x, y) {
		return nil, ErrCanvasDoesNotContainPoint
	}

	return c.Tiles[y][x], nil
}

func (c *Canvas) SetTile(x, y int, t *Tile) error {
	if !c.ContainsPoint(x, y) {
		return ErrCanvasDoesNotContainPoint
	}

	c.markDirty()

	if c.Tiles[y][x] == nil {
		c.Tiles[y][x] = t
		return nil
	}

	if t.Style.bg.IsEmpty() && c.Tiles[y][x].Style != nil {
		t.Style.bg = c.Tiles[y][x].Style.bg
	}
	if t.Style.fg.IsEmpty() && c.Tiles[y][x].Style != nil {
		t.Style.fg = c.Tiles[y][x].Style.fg
	}

	c.Tiles[y][x] = t
	return nil
}

func (c *Canvas) AppendChild(children ...*Canvas) {
	for _, child := range children {
		child.z += c.z
		child.parent = c
		c.Children = append(c.Children, child)
	}
	c.markDirty()
}

func (c *Canvas) RemoveChild(children ...*Canvas) {
	if len(c.Children) == 0 {
		return
	}

	toRemove := make(map[*Canvas]struct{}, len(children))
	toRemain := make([]*Canvas, len(c.Children)-len(children))

	for _, child := range children {
		toRemove[child] = struct{}{}
	}

	for _, child := range c.Children {
		if _, found := toRemove[child]; found {
			continue
		}
		toRemain = append(toRemain, child)
	}
}

// Transfers the tile and children data from another canvas.
// Use this to replace the contents of a canvas while preserving any pointer
// references and transform data.
func (c *Canvas) ReplaceContentsWith(target *Canvas) {
	c.Tiles = append(c.Tiles[:0], target.Tiles...)
	c.Children = append(c.Children[:0], target.Children...)
	c.Width = target.Width
	c.Height = target.Height
	c.markDirty()
	c.markChildrenDirty()
}

// Set the local position - relative to parent.
func (c *Canvas) Position(x, y int) {
	if c.x == x && c.y == y {
		return
	}
	c.x = x
	c.y = y
	c.markDirty()
}

func (c *Canvas) GetPosition() (int, int) {
	return c.x, c.y
}

// Set the public position. The transform pointer should be shared with a
// control if they're to be connected.
func (c *Canvas) Transform(x, y int) {
	c.transform.x += x
	c.transform.y += y
	c.markDirty()
	c.markChildrenDirty()
}

func (c *Canvas) GetTransform() (int, int) {
	return c.transform.x, c.transform.y
}

func (c *Canvas) SetTransform(x, y int) {
	c.transform.x = x
	c.transform.y = y
	c.markDirty()
	c.markChildrenDirty()
}

// A canvas will use a locally bound transform unless otherwise specified.
// Assign a new one if the transform must be shared elsewhere.
func (c *Canvas) AssignTransform(t *Transform) {
	c.transform = t
	c.markDirty()
	c.markChildrenDirty()
}

func (c Canvas) ContainsPoint(x, y int) bool {
	return x >= 0 &&
		x < c.Width &&
		y >= 0 &&
		y < c.Height
}

func (c *Canvas) SetSize(width, height int) {
	tiles := make([][]*Tile, height)
	for i := range tiles {
		tiles[i] = make([]*Tile, width)
	}
}

func (c *Canvas) FlushChildren() {
	c.Children = c.Children[:0]
	c.markDirty()
}

func (c *Canvas) LoopTiles(f func(x, y int, tile *Tile)) {
	for y := range c.Tiles {
		for x := range c.Tiles[y] {
			f(x, y, c.Tiles[y][x])
		}
	}
}

func (c Canvas) X() int {
	return c.x
}

func (c Canvas) Y() int {
	return c.y
}

func (c Canvas) Z() int {
	return c.z
}

func (c *Canvas) SetZ(z int) {
	c.z = z
}
