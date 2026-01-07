package tengin

import (
	"cmp"
	"slices"

	"github.com/gdamore/tcell/v3"
)

// A scene is used by the engine to render canvases. Only one scene should be
// provided to the renderer
type Scene struct {
	canvases       []*Canvas
	controlManager *controlManager
	layers         map[*Canvas]*layer
	cachedLayers   []*layer
	cellBuffer     [][]*cell
	defaultStyle   *Style
	debugOps       []*drawOp
	dirtyZ         bool
	screenRect     Rect
}

type layer struct {
	z       int
	drawOps []*drawOp
	root    *Canvas
	dirty   bool
	dirtyZ  bool
}

type drawOp struct {
	x, y int
	z    int
	tile *Tile
}

type cell struct {
	char string
	bg   Color
	fg   Color
}

func NewScene(width, height int) *Scene {
	s := &Scene{
		canvases:       []*Canvas{},
		controlManager: newControlManager(),
		defaultStyle:   NewStyle(),
		layers:         map[*Canvas]*layer{},
		cachedLayers:   []*layer{},
		dirtyZ:         true,
		screenRect:     NewRect(0, 0, width-1, height-1),
		debugOps:       []*drawOp{},
	}

	s.cellBuffer = make([][]*cell, height)
	for y := range s.cellBuffer {
		s.cellBuffer[y] = make([]*cell, width)

		for x := range s.cellBuffer[y] {
			s.cellBuffer[y][x] = newCell(s.defaultStyle.bg, s.defaultStyle.fg)
		}
	}

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

func newCell(bg, fg Color) *cell {
	return &cell{
		char: "",
		bg:   bg,
		fg:   fg,
	}
}

func (s *Scene) AppendCanvas(c ...*Canvas) {
	for _, canvas := range c {
		s.canvases = append(s.canvases, canvas)
	}
}

func (s *Scene) RemoveCanvas(c ...*Canvas) {
	if len(s.canvases) == 0 {
		return
	}

	toRemove := make(map[*Canvas]struct{}, len(c))
	toRemain := make([]*Canvas, len(s.canvases)-len(c))

	for _, canvas := range c {
		toRemove[canvas] = struct{}{}
	}

	for _, canvas := range s.canvases {
		if _, found := toRemove[canvas]; found {
			continue
		}
		toRemain = append(toRemain, canvas)
	}
}

func (s *Scene) AppendControl(c ...*Control) {
	cm := s.controlManager
	for _, ctrl := range c {
		cm.AppendControl(ctrl)
	}
}

func (s *Scene) RemoveControl(c ...*Control) {
	s.controlManager.RemoveControl(c...)
}

func (s *Scene) HitTest(x, y int) *Control {
	cm := s.controlManager
	for i := len(cm.controls) - 1; i >= 0; i-- {
		toMatch := cm.controls[i]
		if !toMatch.ContainsPoint(x, y) {
			continue
		}

		return toMatch
	}

	return nil
}

func (s *Scene) SetDefaultStyle(def *Style) {
	s.defaultStyle = def
}

func (s *Scene) OnScreenResize(width, height int) {
	s.screenRect = NewRect(0, 0, width-1, height-1)
}

//func (s *Scene) Flush() {
//	s.canvases = s.canvases[:0]
//	s.layers = map[*Canvas]*layer{}
//}

func (s *Scene) update() {
	if s == nil {
		return
	}

	if s.controlManager.IsDirty() {
		s.controlManager.Sort()
	}
}

var (
	renderProfilerLayers    = NewDebugTimer("Layers")
	renderProfilerClrBuffer = NewDebugTimer("Clr Buffer")
	renderProfilerCompose   = NewDebugTimer("Compose")
	renderProfilerSort      = NewDebugTimer("Sort")
	renderProfilerRender    = NewDebugTimer("Render")
)

func (s *Scene) render(screen tcell.Screen, debug *Canvas) {
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

	for y := range s.cellBuffer {
		for x := range s.cellBuffer[y] {
			s.cellBuffer[y][x] = newCell(s.defaultStyle.bg, s.defaultStyle.fg)
		}
	}

	allDrawOps := 0
	renderedDrawOps := 0
	for _, layer := range s.cachedLayers {
		for _, op := range layer.drawOps {
			allDrawOps++
			if !s.screenRect.Contains(op.x, op.y) || op.tile == nil {
				continue
			}
			renderedDrawOps++

			style := op.tile.Style
			cell := s.cellBuffer[op.y][op.x]

			if !style.bg.IsEmpty() {
				cell.bg = style.bg
			}
			if !style.fg.IsEmpty() {
				cell.fg = style.fg
			}
			if op.tile.Char != "" {
				cell.char = op.tile.Char
			}
		}
	}

	for y := range s.cellBuffer {
		for x, cell := range s.cellBuffer[y] {
			style := tcell.StyleDefault.
				Background(cell.bg.tcell()).
				Foreground(cell.fg.tcell())

			screen.Put(x, y, string(cell.char), style)
		}
	}

	renderProfilerRender.End()

	s.debugOps = s.debugOps[:0]
	debug.compose(&s.debugOps)
	for _, op := range s.debugOps {
		style := tcell.StyleDefault.
			Background(op.tile.Style.bg.tcell()).
			Foreground(op.tile.Style.fg.tcell())
		screen.Put(op.x, op.y, op.tile.Char, style)
	}

	DebugLog("Draw Ops - All", allDrawOps)
	DebugLog("Draw Ops - Render", renderedDrawOps)

	// s.flush()
}
