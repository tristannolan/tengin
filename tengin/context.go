package tengin

import "github.com/tristannolan/tengin/tengin/internal/systems"

type Context struct {
	Debug  *systems.Debugger
	Input  *InputService
	Render *RenderService
}
