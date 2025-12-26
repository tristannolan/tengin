package tengin

import "math"

type Context interface {
	// Input
	Key() Key
	LastKey() Key
	MouseKey() Mouse
	LastMouseKey() Mouse
	MouseWheel() Mouse
	LastMouseWheel() Mouse

	// Engine
	Tick() int
	TickRate() float64
	ScreenSize() (int, int)
	ScreenResizing() bool
	ScreenFocused() bool
	Quit()

	// Drawing
	SubmitScene(s *Scene)
}

type frameContext struct {
	e *Engine
}

func newFrameContext(e *Engine) frameContext {
	return frameContext{
		e: e,
	}
}

func (c frameContext) Key() Key {
	return c.e.input.key
}

func (c frameContext) LastKey() Key {
	return c.e.input.lastKey
}

func (c frameContext) MouseKey() Mouse {
	return c.e.input.mouseKey
}

func (c frameContext) LastMouseKey() Mouse {
	return c.e.input.lastMouseKey
}

func (c frameContext) MouseWheel() Mouse {
	return c.e.input.mouseWheel
}

func (c frameContext) LastMouseWheel() Mouse {
	return c.e.input.lastMouseWheel
}

func (c frameContext) Tick() int {
	return c.e.getTick()
}

func (c frameContext) TickRate() float64 {
	return math.Round(c.e.getTickRate())
}

func (c frameContext) Quit() {
	c.e.stopRunning()
}

func (c frameContext) ScreenResizing() bool {
	return c.e.input.isScreenResizing
}

func (c frameContext) ScreenFocused() bool {
	return c.e.input.isScreenFocused
}

func (c frameContext) SubmitScene(s *Scene) {
	c.e.scene = s
}

func (c frameContext) ScreenSize() (int, int) {
	return c.e.ScreenSize()
}
