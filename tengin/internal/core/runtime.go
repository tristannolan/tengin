package core

type Runtime struct {
	tick               int
	deltaTime          float64
	tps                int
	fps                int
	updateAccummulator float64
	drawAccummulator   float64
}
