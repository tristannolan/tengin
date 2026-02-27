// The tengin package provides the tools needed to create an interactive,
// tick based program in the terminal.
package tengin

import (
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

func New(configs ...Config) (*Engine, error) {
	e := &Engine{
		Config:    NewDefaultConfig(),
		Lifecycle: core.NewLifecycle(),
		Runtime:   core.NewRuntime(),
	}

	for _, config := range configs {
		e.LoadConfig(config)
	}

	linkDebug(e)

	return e, nil
}

func (e *Engine) LoadConfig(c Config) {}

func (e *Engine) Stop() {
	func() {
		// e.screen.Fini()
		if r := recover(); r != nil {
			panic(r)
		}
	}()
}

func (e *Engine) Run(d Driver) error {
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
		// sleep
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
	// handle pause/unpause
	if !e.Lifecycle.ShouldUpdate() {
		return
	}
	// advance tick
	// generate context
	// d.Update()
}

func (e *Engine) draw(d Driver) {
	if !e.Lifecycle.ShouldUpdate() {
		return
	}
	// d.Draw()
	// e.Systems.Render()
}

// loop control

func (e *Engine) loopTimingCalc()    // edit runtime values
func (e *Engine) updateCycles() bool // updates remaining in loop cycle
func (e *Engine) drawCycles() bool   // draws remaining in loop cycle
