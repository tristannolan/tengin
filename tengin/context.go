package tengin

import "github.com/gdamore/tcell/v3"

type Context interface {
	// Input
	Key() tcell.Key
	Str() string
	// PutStr(x, y int, str string)

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

func (c frameContext) Str() string {
	return c.e.input.Str()
}

func (c frameContext) Key() tcell.Key {
	return c.e.input.Key()
}

func (c frameContext) Tick() int {
	return c.e.getTick()
}

func (c frameContext) Quit() {
	c.e.stopRunning()
}

//func (c frameContext) PutStr(x, y int, str string) {
//	c.e.screen.PutStr(x, y, str)
//}

func (c frameContext) NewScene() Scene {
	return newScene()
}

func (c frameContext) SubmitScene(s Scene) {
	c.e.scene = s
}
