package tengin

import "github.com/tristannolan/tengin/tengin/internal/systems"

type Services struct {
	Input  *InputService
	Render *RenderService
	// Audio  *AudioService
	TermDriver *TermService
}

func NewServices() (*Services, error) {
	termDriver, err := systems.NewTcellTermDriver()
	if err != nil {
		return nil, err
	}

	s := &Services{
		Render:     NewRenderService(termDriver),
		Input:      NewInputService(termDriver),
		TermDriver: NewTermService(termDriver),
	}

	return s, nil
}
