package tengin

import (
	"math"
	"sync"
	"time"

	"github.com/gdamore/tcell/v3"
)

type Game interface {
	Update(ctx Context)
	Draw(ctx Context)
}

type Engine struct {
	mu                sync.RWMutex
	input             input
	liveInput         *liveInput
	screen            tcell.Screen
	scene             *Scene
	debug             *debug
	running           bool
	tick              int
	tickRate          float64
	frameRate         float64
	tps               int
	fps               int
	deltaTime         float32
	runWhenUnfocused  bool
	drawWhenUnfocused bool
	paused            bool
	shouldAdvanceTick bool
	advanceTickTarget int
	advanceTickCount  int
}

func New() (*Engine, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}
	if err := screen.Init(); err != nil {
		return nil, err
	}

	screen.EnableMouse()
	screen.EnableFocus()
	screen.SetTitle("Tengin")

	w, h := screen.Size()

	e := &Engine{
		mu:                sync.RWMutex{},
		input:             newInput(),
		liveInput:         newLiveInput(),
		screen:            screen,
		running:           true,
		tick:              0,
		tickRate:          60,
		frameRate:         20,
		tps:               0,
		fps:               0,
		debug:             newDebug(w, h),
		deltaTime:         1,
		runWhenUnfocused:  true,
		drawWhenUnfocused: true,
		shouldAdvanceTick: false,
		advanceTickTarget: 1,
		advanceTickCount:  0,
	}

	e.debug.registerCommands(e)

	return e, nil
}

func (e *Engine) Run(g Game) error {
	ctx := newFrameContext(e)

	lastTime := time.Now()
	lastStatTime := lastTime

	updateAcc := 0.0
	drawAcc := 0.0

	tickCount := 0
	frameCount := 0

	updateProfiler := NewDebugTimer("Update Cycle")
	drawProfiler := NewDebugTimer("Draw Cycle")

	e.liveInput.listen(e.screen)

	for e.isRunning() {
		now := time.Now()
		dt := now.Sub(lastTime).Seconds()
		lastTime = now

		updateAcc += dt
		drawAcc += dt

		tickDur := 1.0 / e.tickRate
		frameDur := 1.0 / e.frameRate

		for updateAcc >= tickDur {
			updateAcc -= tickDur

			updateProfiler.Start()
			Update(e, g, &ctx)
			updateProfiler.End()

			tickCount++
		}

		if drawAcc > frameDur {
			drawAcc = 0
			drawProfiler.Start()
			Draw(e, g, &ctx)
			drawProfiler.End()

			frameCount++
		}

		if now.Sub(lastStatTime).Seconds() >= 1.0 {
			e.tps = tickCount
			e.fps = frameCount
			tickCount = 0
			frameCount = 0
			lastStatTime = now
		}

		minDur := math.Min(tickDur-updateAcc, frameDur-drawAcc)
		if minDur > 0.001 {
			time.Sleep(time.Duration(minDur * float64(time.Second)))
		} else {
			time.Sleep(time.Millisecond)
		}
	}

	return nil
}

func Update(e *Engine, g Game, ctx *frameContext) {
	if e.runWhenUnfocused && !e.isScreenFocused() {
		return
	}
	e.input.poll(e.liveInput)

	e.debug.handleCommandInput(ctx.Key())

	if e.isScreenResizing() == true {
		e.syncScreenSize()
	}

	if e.paused {
		switch ctx.Key().Value() {
		case "f":
			e.advanceTick(1)
		case "F":
			e.advanceTick(10)
		}
	}
	if e.shouldAdvanceTick {
		if e.advanceTickCount < e.advanceTickTarget {
			e.advanceTickCount++
		} else {
			e.advanceTickCount = 0
			e.shouldAdvanceTick = false
		}
	}
	if e.isPaused() {
		return
	}

	e.incrementTick()
	// Update
	g.Update(ctx)
	e.scene.update(ctx.Key())
}

var (
	profilerShowScreen = NewDebugTimer("Tcell.Show()")
	profilerGameDraw   = NewDebugTimer("Game.Draw()")
	profilerRender     = NewDebugTimer("Scene.Render()")
)

func Draw(e *Engine, g Game, ctx *frameContext) {
	if e.drawWhenUnfocused && !e.isScreenFocused() {
		return
	}
	DebugLog("Input", e.input.lastKey.Value())
	DebugLog("Tick", e.getTick())
	DebugLog("TPS", e.tps)
	DebugLog("FPS", e.fps)

	profilerGameDraw.Start()

	g.Draw(ctx)

	profilerGameDraw.End()
	profilerRender.Start()

	// e.screen.Clear()

	e.debug.draw(e.screen)
	e.scene.render(e.screen, e.debug.canvas)

	profilerRender.End()
	profilerShowScreen.Start()

	e.screen.Show()

	profilerShowScreen.End()
}

func (e *Engine) Quit() {
	func() {
		e.screen.Fini()
		if r := recover(); r != nil {
			panic(r)
		}
	}()
}

func (e *Engine) ScreenSize() (int, int) {
	return e.screen.Size()
}

func (e *Engine) SetDefaultStyle(s Style) {
	e.screen.SetStyle(tcell.StyleDefault.
		Background(s.bg.tcell()).
		Foreground(s.fg.tcell()),
	)
}

func (e *Engine) SetTickRate(i int) {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.tickRate = float64(i)
}

func (e *Engine) SetFrameRate(i int) {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.frameRate = float64(i)
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

func (e *Engine) getTickRate() float64 {
	e.mu.RLock()
	defer e.mu.RUnlock()

	return e.tickRate
}

func (e *Engine) incrementTick() {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.tick++
}

func (e *Engine) syncScreenSize() {
	e.screen.Sync()
	e.liveInput.onScreenResizeComplete()
}

func (e *Engine) isScreenResizing() bool {
	return e.input.isScreenResizing
}

func (e *Engine) isScreenFocused() bool {
	return e.input.isScreenFocused
}

func (e *Engine) isPaused() bool {
	if e.paused && e.shouldAdvanceTick {
		return false
	}
	return e.paused
}

func (e *Engine) pause() {
	e.paused = true
}

func (e *Engine) unpause() {
	e.paused = false
}

func (e *Engine) advanceTick(i int) {
	e.shouldAdvanceTick = true
	e.advanceTickTarget = i
	e.advanceTickCount = 0
	ConsoleLogF("Advance Tick: %d", i)
}

func (e *Engine) SetDebug(b bool) {
	sw, sh := e.ScreenSize()
	if b == true {
		e.debug.Enable(sw, sh)
	} else {
		e.debug.Disable()
	}
}
