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
	mu       sync.RWMutex
	input    input
	screen   tcell.Screen
	running  bool
	tick     int
	tickRate int // Ticks per second
	debug    debug
	scene    Scene
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
		mu:       sync.RWMutex{},
		input:    newInput(),
		screen:   screen,
		running:  true,
		tick:     0,
		tickRate: 60,
		debug:    newDebug(),
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

	ticker := time.NewTicker(time.Second / time.Duration(e.tickRate))
	defer ticker.Stop()

	// Main game loop
	for e.isRunning() {
		i := &e.input

		// Engine management
		e.incrementTick()
		i.poll()

		if e.input.isResizingScreen == true {
			e.syncScreenSize()
		}

		e.debug.update()
		if i.lastKey().isRuneKey() {
			DebugLog("Input", i.lastKey().rune)
		} else if i.lastKey().isSpecialKey() {
			DebugLog("Input", int(i.lastKey().special))
		}
		DebugLog("Key Empty", ctx.KeyIsEmpty())
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

func (e *Engine) SetTickRate(i int) {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.tickRate = i
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
