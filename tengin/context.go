package tengin

type Context interface {
	// Input
	Key() Key
	KeyIsRune(r rune) bool
	KeyIsSpecial(k specialKey) bool

	// Engine
	Tick() int
	Quit()

	// Drawing
	NewScene() Scene
	SubmitScene(s Scene)
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
	return c.e.input.Key()
}

func (c frameContext) KeyIsRune(r rune) bool {
	return c.e.input.Key().kind == keyRune && c.e.input.Key().rune == r
}

func (c frameContext) KeyIsSpecial(k specialKey) bool {
	return c.e.input.Key().kind == keySpecial && c.e.input.Key().special == k
}

func (c frameContext) Tick() int {
	return c.e.getTick()
}

func (c frameContext) Quit() {
	c.e.stopRunning()
}

func (c frameContext) NewScene() Scene {
	return newScene()
}

func (c frameContext) SubmitScene(s Scene) {
	c.e.scene = s
}
