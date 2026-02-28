package tengin

type RenderServicer interface {
	Show()
	SetTile(x, y int, ch string)
	Size() (w, h int)
}

type RenderService struct {
	service RenderServicer
}

func NewRenderService(r RenderServicer) *RenderService {
	s := &RenderService{
		service: r,
	}

	return s
}

func (s *RenderService) Show() {
	s.service.Show()
}
