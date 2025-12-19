package tengin

import (
	"sync"
	"time"

	"github.com/gdamore/tcell/v3"
)

// Implemented by the user
// Calls functions to be presented in the game loop
type Game interface {
	Update(ctx Context)
	Draw(ctx Context)
}

// [Loop, Ticks, input, Drawing]
type Engine struct {
	mu      sync.RWMutex
	input   input
	screen  tcell.Screen
	running bool
	tick    int
	debug   debug
}

func New() (*Engine, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}
	if err := screen.Init(); err != nil {
		return nil, err
	}

	e := &Engine{
		mu:      sync.RWMutex{},
		input:   newInput(),
		screen:  screen,
		running: true,
		tick:    0,
		debug:   newDebug(),
	}

	return e, nil
}

// Runs the basic game loop
func (e *Engine) Run(g Game) error {
	defer func() {
		e.screen.Fini()
		if r := recover(); r != nil {
			panic(r)
		}
	}()

	ctx := newFrameContext(e)

	e.input.listen(e)

	ticker := time.NewTicker(time.Second / time.Duration(60))
	defer ticker.Stop()

	// Main game loop
	for e.isRunning() {
		// Update
		e.incrementTick()

		e.input.poll()
		e.debug.update()
		g.Update(ctx)

		// Draw
		e.screen.Clear()

		g.Draw(ctx)

		e.debug.log("Input", e.input.Str())
		e.debug.log("Tick", e.getTick())
		e.debug.draw(e.screen)
		e.screen.Show()

		<-ticker.C
	}

	return nil
}

func (e *Engine) isRunning() bool {
	e.mu.RLock()
	defer e.mu.RUnlock()

	return e.running
}

func (e *Engine) stopRunning() {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.running = false
}

func (e *Engine) getTick() int {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.tick
}

func (e *Engine) incrementTick() {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.tick++
}
