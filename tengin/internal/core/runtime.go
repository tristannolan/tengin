package core

type Runtime struct {
	Tick               int
	DeltaTime          float64
	TPS                int
	FPS                int
	UpdateAccummulator float64
	DrawAccummulator   float64
}

func NewRuntime() *Runtime {
	r := &Runtime{
		Tick:               0,
		DeltaTime:          0.0,
		TPS:                0,
		FPS:                0,
		UpdateAccummulator: 0.0,
		DrawAccummulator:   0.0,
	}

	return r
}
