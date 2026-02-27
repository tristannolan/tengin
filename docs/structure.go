package docs

type Engine struct {
	Config    *Config
	Runtime   *Runtime
	Lifecycle *Lifecycle
	Systems   *Systems
	Services  *Services
}

type Driver interface {
	Update(ctx *Context)
	Draw(ctx *Context)
}

// Customise the engine configuration
type Config struct {
	Title     string
	TickRate  float64
	FrameRate float64
}

// Handle shutdown, pause, restart, etc
type Lifecycle struct {
	running           bool
	paused            bool
	shutdownRequested bool
}

// Everything that changes frame to frame
// This is a data only structure, and should not be doing calculations. Perform
// any operations in a system, and only get/set with runtime
type Runtime struct {
	tick               int
	deltaTime          float64
	tps                int
	fps                int
	updateAccummulator float64
	drawAccummulator   float64
}

// Internal tools to make it all happen
// These are the interface points with everything that the engine can do. These
// should never be touched by the user directly
// Engine control should happen through core functions (new|init|update|shutdown)
// User control via services has the fluid structure
type Systems struct {
	Debug  *DebugSystem
	Input  *InputSystem
	Audio  *AudioSystem
	Render *RenderSystem
	Screen *ScreenSystem
}

type (
	DebugSystem  struct{} // file logging, console logging
	InputSystem  struct{} // holds tcell input listener
	AudioSystem  struct{}
	RenderSystem struct{} // holds logic to draw to screen
	ScreenSystem struct{} // holds tcell
)

// External interface to use internal systems
// It is purely a way to reach into systems, and should do no real work itself.
// That said, the service should also not blindly expose the system.
// Each service should handle only one system.
type Services struct {
	Debug  *DebugService
	Input  *InputService
	Audio  *AudioService
	Render *RenderService
}

type DebugService struct {
	system *DebugSystem
}
type InputService struct {
	system *InputSystem
}
type AudioService struct {
	system *AudioSystem
}
type RenderService struct {
	system *RenderSystem
}

// ===========
//
//	Context
//
// ===========
type Context struct {
	Debug  *DebugService
	Input  *InputService
	Render *RenderService
}

// =============
//
//	Interface
//
// =============
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
