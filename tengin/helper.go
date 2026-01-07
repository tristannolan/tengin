package tengin

type Rect struct {
	minX, minY int
	maxX, maxY int
}

func NewRect(minX, minY, maxX, maxY int) Rect {
	return Rect{minX, minY, maxX, maxY}
}

func (r Rect) Contains(x, y int) bool {
	return x >= r.minX &&
		x <= r.maxX &&
		y >= r.minY &&
		y <= r.maxY
}

type Transform struct {
	x, y int
}

func NewTransform(x, y int) *Transform {
	return &Transform{
		x: x,
		y: y,
	}
}

func (t Transform) X() int {
	return t.x
}

func (t Transform) Y() int {
	return t.y
}
