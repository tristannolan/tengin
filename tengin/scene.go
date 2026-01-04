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
	layers       []*layer
}

type drawOp struct {
	x, y int
	z    int
	tile *Tile
}

type layer struct {
	z       int
	drawOps []*drawOp
}

func NewScene() *Scene {
	return &Scene{
		canvases:     []*Canvas{},
		defaultStyle: NewStyle(),
		layers:       []*layer{},
	}
}

func newDrawOp(x, y, z int, tile *Tile) *drawOp {
	return &drawOp{
		x:    x,
		y:    y,
		z:    z,
		tile: tile,
	}
}

func newLayer(z int) layer {
	return layer{
		z:       z,
		drawOps: []*drawOp{},
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
)

func (s *Scene) newRender(screen tcell.Screen) {
	renderProfilerCompose.Start()
	layers := []*layer{}
	for _, c := range s.canvases {
		layer := newLayer(c.Z)
		c.compose(&layer.drawOps)

		layers = append(layers, &layer)
	}
	renderProfilerCompose.End()

	// Sort local z axis for each canvas layer
	for i := range layers {
		slices.SortStableFunc(layers[i].drawOps, func(a, b *drawOp) int {
			return cmp.Compare(a.z, b.z)
		})
	}

	// Now sort the layers
	slices.SortStableFunc(layers, func(a, b *layer) int {
		return cmp.Compare(a.z, b.z)
	})

	// We can ignore any tiles that aren't in the screen
	screenWidth, screenHeight := screen.Size()
	clip := NewRect(0, 0, screenWidth-1, screenHeight-1)

	// Store the colour for each tile because tcell is silly.
	// There's no such thing as transparency (sometimes...), so foreground
	// only styles will output the terminal default as a background.
	// We need to bubble the colours up to avoid this.
	bgBuffer := make([][]Color, screenHeight)
	fgBuffer := make([][]Color, screenHeight)
	for y := range bgBuffer {
		bgBuffer[y] = make([]Color, screenWidth)
		fgBuffer[y] = make([]Color, screenWidth)

		for x := range bgBuffer[y] {
			bgBuffer[y][x] = s.defaultStyle.bg
			fgBuffer[y][x] = s.defaultStyle.fg
		}
	}

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
	renderProfilerRender.Start()
	for i := range layers {
		for _, op := range layers[i].drawOps {
			allDrawOps++
			if !clip.Contains(op.x, op.y) || op.tile == nil {
				continue
			}
			renderedDrawOps++

			// Check if drawOp has different background to buffer
			bgColor := op.tile.Style.bg
			fgColor := op.tile.Style.fg
			if !bgColor.IsEmpty() {
				bgBuffer[op.y][op.x] = bgColor
			}
			if !fgColor.IsEmpty() {
				fgBuffer[op.y][op.x] = fgColor
			}

			style := tcell.StyleDefault
			style = style.Background(bgBuffer[op.y][op.x].tcell())
			style = style.Foreground(fgBuffer[op.y][op.x].tcell())

			if !fgColor.IsEmpty() {
				style = style.Foreground(fgColor.tcell())
			}

			screen.Put(op.x, op.y, string(op.tile.Char), style)
		}
	}
	renderProfilerRender.End()

	DebugLog("Draw Ops - All", allDrawOps)
	DebugLog("Draw Ops - Render", renderedDrawOps)

	s.flush()
}

func (s *Scene) render(screen tcell.Screen) {
	layers := []*layer{}
	for _, c := range s.canvases {
		layer := newLayer(c.Z)
		c.compose(&layer.drawOps)

		layers = append(layers, &layer)
	}

	// Sort local z axis for each canvas layer
	for i := range layers {
		slices.SortStableFunc(layers[i].drawOps, func(a, b *drawOp) int {
			return cmp.Compare(a.z, b.z)
		})
	}

	// Now sort the layers
	slices.SortStableFunc(layers, func(a, b *layer) int {
		return cmp.Compare(a.z, b.z)
	})

	// We can ignore any tiles that aren't in the screen
	screenWidth, screenHeight := screen.Size()
	clip := NewRect(0, 0, screenWidth-1, screenHeight-1)

	// Store the colour for each tile because tcell is silly.
	// There's no such thing as transparency (sometimes...), so foreground
	// only styles will output the terminal default as a background.
	// We need to bubble the colours up to avoid this.
	bgBuffer := make([][]Color, screenHeight)
	fgBuffer := make([][]Color, screenHeight)
	for y := range bgBuffer {
		bgBuffer[y] = make([]Color, screenWidth)
		fgBuffer[y] = make([]Color, screenWidth)

		for x := range bgBuffer[y] {
			bgBuffer[y][x] = s.defaultStyle.bg
			fgBuffer[y][x] = s.defaultStyle.fg
		}
	}

	// Make the scene show the default everywhere
	defStyle := tcell.StyleDefault
	if bg := s.defaultStyle.bg; !bg.IsEmpty() {
		defStyle = defStyle.Background(bg.tcell())
	}
	if fg := s.defaultStyle.fg; !fg.IsEmpty() {
		defStyle = defStyle.Foreground(fg.tcell())
	}
	screen.SetStyle(defStyle)

	for i := range layers {
		for _, op := range layers[i].drawOps {
			if !clip.Contains(op.x, op.y) || op.tile == nil {
				continue
			}

			// Check if drawOp has different background to buffer
			bgColor := op.tile.Style.bg
			fgColor := op.tile.Style.fg
			if !bgColor.IsEmpty() {
				bgBuffer[op.y][op.x] = bgColor
			}
			if !fgColor.IsEmpty() {
				fgBuffer[op.y][op.x] = fgColor
			}

			style := tcell.StyleDefault
			style = style.Background(bgBuffer[op.y][op.x].tcell())
			style = style.Foreground(fgBuffer[op.y][op.x].tcell())

			if !fgColor.IsEmpty() {
				style = style.Foreground(fgColor.tcell())
			}

			screen.Put(op.x, op.y, string(op.tile.Char), style)
		}
	}

	s.flush()
}
