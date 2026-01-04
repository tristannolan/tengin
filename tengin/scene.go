package tengin

import (
	"cmp"
	"slices"

	"github.com/gdamore/tcell/v3"
)

// A scene is used by the engine to render canvases. Only one scene should be
// provided to the renderer
type Scene struct {
	canvases     []*Canvas
	defaultStyle *Style
	layers       map[*Canvas]*layer
	cachedLayers []*layer
	bgBuffer     [][]Color
	fgBuffer     [][]Color
	dirtyZ       bool
}

type drawOp struct {
	x, y int
	z    int
	tile *Tile
}

type layer struct {
	z       int
	drawOps []*drawOp
	root    *Canvas
	dirty   bool
	dirtyZ  bool
}

func NewScene(width, height int) *Scene {
	s := &Scene{
		canvases:     []*Canvas{},
		defaultStyle: NewStyle(),
		layers:       map[*Canvas]*layer{},
		cachedLayers: []*layer{},
		dirtyZ:       true,
	}

	bgBuffer := make([][]Color, height)
	fgBuffer := make([][]Color, height)
	for y := range bgBuffer {
		bgBuffer[y] = make([]Color, width)
		fgBuffer[y] = make([]Color, width)

		for x := range bgBuffer[y] {
			bgBuffer[y][x] = s.defaultStyle.bg
			fgBuffer[y][x] = s.defaultStyle.fg
		}
	}

	s.bgBuffer = bgBuffer
	s.fgBuffer = fgBuffer

	return s
}

func newDrawOp(x, y, z int, tile *Tile) *drawOp {
	return &drawOp{
		x:    x,
		y:    y,
		z:    z,
		tile: tile,
	}
}

func newLayer(z int, c *Canvas) *layer {
	return &layer{
		z:       z,
		drawOps: []*drawOp{},
		root:    c,
		dirty:   true,
		dirtyZ:  true,
	}
}

func (s *Scene) AppendCanvas(c ...*Canvas) {
	for _, canvas := range c {
		s.canvases = append(s.canvases, canvas)
	}
}

func (s *Scene) SetDefaultStyle(def *Style) {
	s.defaultStyle = def
}

func (s *Scene) flush() {
	s.canvases = s.canvases[:0]
}

var (
	renderProfilerLayers    = NewDebugTimer("Layers")
	renderProfilerClrBuffer = NewDebugTimer("Clr Buffer")
	renderProfilerCompose   = NewDebugTimer("Compose")
	renderProfilerRender    = NewDebugTimer("Render")
	renderProfilerSort      = NewDebugTimer("Sort")
)

func (s *Scene) newRender(screen tcell.Screen) {
	renderProfilerCompose.Start()

	for _, c := range s.canvases {
		layer := s.layers[c]

		if layer == nil {
			layer = newLayer(c.Z, c)
			s.layers[c] = layer
			s.cachedLayers = append(s.cachedLayers, layer)
			s.dirtyZ = true
		}

		if c.dirty {
			layer.dirty = true
			c.dirty = false
		}

		if layer.dirty {
			layer.drawOps = layer.drawOps[:0]
			c.compose(&layer.drawOps)
			layer.dirty = false

			slices.SortStableFunc(layer.drawOps, func(a, b *drawOp) int {
				return cmp.Compare(a.z, b.z)
			})
		}

		if layer.dirtyZ {
			layer.z = c.Z
			s.dirtyZ = true
			layer.dirtyZ = false
		}
	}

	renderProfilerCompose.End()
	renderProfilerSort.Start()

	if s.dirtyZ {
		slices.SortStableFunc(s.cachedLayers, func(a, b *layer) int {
			return cmp.Compare(a.z, b.z)
		})

		s.dirtyZ = false
	}

	renderProfilerSort.End()
	renderProfilerRender.Start()

	// We can ignore any tiles that aren't in the screen
	screenWidth, screenHeight := screen.Size()
	clip := NewRect(0, 0, screenWidth-1, screenHeight-1)

	// Make the scene show the default everywhere
	defStyle := tcell.StyleDefault
	if bg := s.defaultStyle.bg; !bg.IsEmpty() {
		defStyle = defStyle.Background(bg.tcell())
	}
	if fg := s.defaultStyle.fg; !fg.IsEmpty() {
		defStyle = defStyle.Foreground(fg.tcell())
	}
	screen.SetStyle(defStyle)

	allDrawOps := 0
	renderedDrawOps := 0
	for _, layer := range s.cachedLayers {
		for _, op := range layer.drawOps {
			allDrawOps++
			if !clip.Contains(op.x, op.y) || op.tile == nil {
				continue
			}
			renderedDrawOps++

			// Check if drawOp has different background to buffer
			bgColor := op.tile.Style.bg
			fgColor := op.tile.Style.fg
			if !bgColor.IsEmpty() {
				s.bgBuffer[op.y][op.x] = bgColor
			}
			if !fgColor.IsEmpty() {
				s.fgBuffer[op.y][op.x] = fgColor
			}

			style := tcell.StyleDefault
			style = style.Background(s.bgBuffer[op.y][op.x].tcell())
			style = style.Foreground(s.fgBuffer[op.y][op.x].tcell())

			if !fgColor.IsEmpty() {
				style = style.Foreground(fgColor.tcell())
			}

			screen.Put(op.x, op.y, string(op.tile.Char), style)
		}
	}
	renderProfilerRender.End()

	DebugLog("Draw Ops - All", allDrawOps)
	DebugLog("Draw Ops - Render", renderedDrawOps)

	// s.flush()
}
