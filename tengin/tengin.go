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

	// Main game loop
	ticker := time.NewTicker(time.Second / time.Duration(60))
	defer ticker.Stop()

	for e.IsRunning() {
		e.input.poll()

		g.Update(ctx)
		g.Draw(ctx)

		<-ticker.C
	}

	return nil
}

func (e *Engine) IsRunning() bool {
	e.mu.RLock()
	defer e.mu.RUnlock()

	return e.running
}

func (e *Engine) StopRunning() {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.running = false
}
