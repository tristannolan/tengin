package tengin

import (
	"github.com/tristannolan/tengin/tengin/internal/core"
	"github.com/tristannolan/tengin/tengin/internal/systems"
)

type Engine struct {
	Config    *Config
	Services  *Services
	Lifecycle *core.Lifecycle
	Systems   *systems.Systems
	Runtime   *core.Runtime
}

type Driver interface {
	Update(ctx *Context)
	Draw(ctx *Context)
}

func New(configs ...Config) *Engine {
	e := Engine{}
	for _, config := range configs {
		e.LoadConfig(config)
	}
	return &Engine{}
}

func (e *Engine) LoadConfig(c Config) {}

func (e *Engine) Run() {
}
