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
