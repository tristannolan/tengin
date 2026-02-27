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
	debugSystem *systems.Debug
}

type Driver interface {
	Update(ctx *Context)
	Draw(ctx *Context)
}

func New(configs ...Config) (*Engine, error) {
	e := Engine{}
	for _, config := range configs {
		e.LoadConfig(config)
	}
	return &Engine{}, nil
}

func (e *Engine) LoadConfig(c Config) {}

func (e *Engine) Run(d Driver) error {
	return nil
}

func (e *Engine) Quit() {
}
