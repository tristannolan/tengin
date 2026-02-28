package tengin

type InputServicer interface {
	Listen()
	Poll()
}

type InputService struct {
	service InputServicer
}

func NewInputService(s InputServicer) *InputService {
	r := &InputService{}

	return r
}

func (s *InputService) Poll() {
	s.service.Poll()
}

func (s *InputService) Listen() {
	s.service.Listen()
}
