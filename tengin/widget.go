package tengin

import "strings"

type Widget interface {
	Canvas() *Canvas
	Control() *Control
}

type Button struct {
	canvas       *Canvas
	control      *Control
	transform    *Transform
	ActiveStyle  *Style
	DefaultStyle *Style
	HoverStyle   *Style
}

func (b Button) Canvas() *Canvas {
	return b.canvas
}

func (b Button) Control() *Control {
	return b.control
}

func NewButton(msg string, def *Style, padding int) Button {
	width := len(msg) + padding*2
	height := 1 + padding*2
	chars := strings.Split(msg, "")

	btn := Button{
		ActiveStyle:  &Style{},
		DefaultStyle: &Style{},
		HoverStyle:   &Style{},
	}
	btn.ActiveStyle.CopyValues(def)
	btn.DefaultStyle.CopyValues(def)
	btn.HoverStyle.CopyValues(def)

	canvas := NewCanvas(width, height)
	canvas.SetAlwaysCache(true)

	for y := range canvas.Tiles {
		for x := range canvas.Tiles[y] {
			char := " "

			if y == padding && x >= padding && x <= len(chars) {
				char = chars[x-padding]
			}

			tile := NewTile(char, btn.ActiveStyle)
			canvas.SetTile(x, y, tile)
		}
	}

	control := NewControl(width, height)
	control.SetHoverAction(func() {
		btn.ActiveStyle = btn.HoverStyle
	})
	control.SetHoverOffAction(func() {
		btn.ActiveStyle = btn.DefaultStyle
	})

	btn.control = control
	btn.canvas = canvas

	btn.AssignTransform(NewTransform(0, 0))

	return btn
}

func (b *Button) AssignTransform(t *Transform) {
	b.transform = t
	b.canvas.AssignTransform(t)
	b.control.AssignTransform(t)
}
