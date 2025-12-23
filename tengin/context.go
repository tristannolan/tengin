package tengin

type Context interface {
	// Input
	KeyIsRune(r rune) bool
	KeyIsSpecial(k specialKey) bool
	KeyIsEmpty() bool
	KeyRuneValue() rune
	KeySpecialValue() specialKey

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

func (c frameContext) Key(r rune) bool {
	key := c.e.input.key()
	return key.isRuneKey() && key.rune == r
}

func (c frameContext) KeyIsRune(r rune) bool {
	key := c.e.input.key()
	return key.isRuneKey() && key.rune == r
}

func (c frameContext) KeyIsSpecial(k specialKey) bool {
	key := c.e.input.key()
	return key.isSpecialKey() && key.special == k
}

func (c frameContext) KeyRuneValue() rune {
	return c.e.input.key().getRuneValue()
}

func (c frameContext) KeySpecialValue() specialKey {
	return c.e.input.key().getSpecialValue()
}

func (c frameContext) KeyIsEmpty() bool {
	key := c.e.input.key()
	return key.isSpecialKey() && key.special == KeyEmpty
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
