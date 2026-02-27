package tengin

import "github.com/tristannolan/tengin/tengin/internal/systems"

type Services struct {
	Debug  *DebugService
	Input  *InputService
	Audio  *AudioService
	Render *RenderService
}

type InputService struct {
	system *systems.Input
}
type AudioService struct {
	system *systems.Audio
}
type RenderService struct {
	system *systems.Render
}
