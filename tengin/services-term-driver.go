package tengin

type TermServicer interface {
	Init() error
	Stop()
}

type TermService struct {
	service TermServicer
}

func NewTermService(r TermServicer) *TermService {
	t := &TermService{
		service: r,
	}

	return t
}

func (s TermService) Init() error {
	return s.Init()
}

func (s TermService) Stop() {
	s.Stop()
}
