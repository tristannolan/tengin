package tengin

import (
	"sync"
	"time"

	"github.com/gdamore/tcell/v3"
)

type Game interface {
	Update(ctx Context)
	Draw(ctx Context)
}

type Engine struct {
	mu      sync.RWMutex
	input   input
	screen  tcell.Screen
	running bool
	tick    int
	debug   debug
	scene   Scene
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

func (e *Engine) Quit() {
	func() {
		e.screen.Fini()
		if r := recover(); r != nil {
			panic(r)
		}
	}()
}

// Runs the basic game loop
func (e *Engine) Run(g Game) error {
	ctx := newFrameContext(e)

	e.input.listen(e.screen)

	ticker := time.NewTicker(time.Second / time.Duration(60))
	defer ticker.Stop()

	// Main game loop
	for e.isRunning() {
		// Engine management
		e.incrementTick()
		e.input.poll()

		if e.input.isResizingScreen == true {
			e.syncScreenSize()
		}

		e.debug.update()
		if e.input.Key().kind == keyRune {
			DebugLog("Input", e.input.Key().rune)
		} else if e.input.Key().kind == keySpecial {
			DebugLog("Input", int(e.input.Key().special))
		}
		DebugLog("Tick", e.getTick())

		// Update
		g.Update(ctx)

		// Draw
		e.screen.Clear()
		g.Draw(ctx)
		e.scene.render(e.screen)
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

func (e *Engine) syncScreenSize() {
	e.screen.Sync()
	e.input.onScreenResizeComplete()
}
