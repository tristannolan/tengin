package tengin

import "github.com/gdamore/tcell/v3"

type Context interface {
	Key() tcell.Key
	Str() string

	Clear()
	Show()
	PutStr(x, y int, str string)

	Quit()
}

type frameContext struct {
	e *Engine
}

func (c frameContext) Str() string {
	return c.e.input.Str()
}

func (c frameContext) Key() tcell.Key {
	return c.e.input.Key()
}

func (c frameContext) Quit() {
	c.e.StopRunning()
}

func (c frameContext) Clear() {
	c.e.screen.Clear()
}

func (c frameContext) Show() {
	c.e.screen.Show()
}

func (c frameContext) PutStr(x, y int, str string) {
	c.e.screen.PutStr(x, y, str)
}

func newFrameContext(e *Engine) frameContext {
	return frameContext{
		e,
	}
}
