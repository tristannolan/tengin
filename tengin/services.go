package tengin

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
