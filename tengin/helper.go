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

type Position struct {
	x, y int
}

type Transform struct {
	x, y int
}

func NewPosition(x, y int) Position {
	return Position{
		x: x,
		y: y,
	}
}

func (p Position) X() int {
	return p.x
}

func (p Position) Y() int {
	return p.y
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

// A transform shouldn't be set directly if the goal is to update a
// canvas+control pair. Reason being that both require being marked as dirty
// to update correctly. The update could be forced, but it's safer to always
// alter the canvas transform and let control tag along for the ride
