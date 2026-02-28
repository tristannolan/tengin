// The tengin package provides the tools needed to create an interactive,
// tick based program in the terminal.
package tengin

import (
	"time"

	"github.com/tristannolan/tengin/tengin/internal/core"
	"github.com/tristannolan/tengin/tengin/internal/systems"
)

type Engine struct {
	Config      *Config
	Services    *Services
	Lifecycle   *core.Lifecycle
	Runtime     *core.Runtime
	Systems     *systems.Systems
	debugSystem *systems.Debugger
}

type Driver interface {
	Update(ctx *Context)
	Draw(ctx *Context)
}

func New() (*Engine, error) {
	services, err := NewServices()
	if err != nil {
		return nil, err
	}

	e := &Engine{
		Config:    NewDefaultConfig(),
		Services:  services,
		Lifecycle: core.NewLifecycle(),
		Runtime:   core.NewRuntime(),
		Systems:   systems.NewSystems(),
	}

	linkDebug(e)

	return e, nil
}

func (e *Engine) LoadConfig(c Config) {}

func (e *Engine) Stop() {
	func() {
		e.Services.TermDriver.Stop()
		if r := recover(); r != nil {
			panic(r)
		}
	}()
}

func (e *Engine) Run(d Driver) error {
	err := e.Services.TermDriver.Init()
	if err != nil {
		return err
	}

	e.Services.Input.Listen()

	e.Lifecycle.Run()

	// loop
	for e.Lifecycle.Running() {
		// start loop timing calculations

		for e.updateCycles() {
			e.update(d)
		}

		for e.drawCycles() {
			e.draw(d)
		}

		// end loop timing calculations
		// calculate sleep duration for new cycle
		time.Sleep(time.Millisecond)

		// e.Lifecycle.RequestShutdown()
	}

	e.Lifecycle.Shutdown()

	return nil
}

// ============
//
//	Internal
//
// ============

func (e *Engine) update(d Driver) {
	// resize interfaces
	// poll input
	e.Services.Input.Poll()
	// handle pause/unpause
	if !e.Lifecycle.ShouldUpdate() {
		return
	}
	// advance tick
	// generate context
	// d.Update()
}

func (e *Engine) draw(d Driver) {
	if !e.Lifecycle.ShouldDraw() {
		return
	}
	// d.Draw()
	e.Services.Render.Show()
}

// loop control

func (e *Engine) updateCycles() bool { // updates remaining in loop cycle
	return false
}

func (e *Engine) drawCycles() bool { // draws remaining in loop cycle
	return false
}
