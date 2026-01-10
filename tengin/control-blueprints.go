package tengin

import "strings"

type button struct {
	Canvas       *Canvas
	Control      *Control
	Transform    *Transform
	ActiveStyle  *Style
	DefaultStyle *Style
	HoverStyle   *Style
}

func NewButton(msg string, def *Style, padding int) button {
	width := len(msg) + padding*2
	height := 1 + padding*2
	chars := strings.Split(msg, "")

	button := button{
		ActiveStyle:  &Style{},
		DefaultStyle: &Style{},
		HoverStyle:   &Style{},
	}
	button.ActiveStyle.CopyValues(def)
	button.DefaultStyle.CopyValues(def)
	button.HoverStyle.CopyValues(def)

	canvas := NewCanvas(width, height)
	canvas.SetAlwaysCache(true)

	for y := range canvas.Tiles {
		for x := range canvas.Tiles[y] {
			char := " "

			if y == padding && x >= padding && x <= len(chars) {
				char = chars[x-padding]
			}

			tile := NewTile(char, button.ActiveStyle)
			canvas.SetTile(x, y, tile)
		}
	}

	control := NewControl(width, height)
	control.SetHoverAction(func() {
		button.ActiveStyle = button.HoverStyle
	})
	control.SetHoverOffAction(func() {
		button.ActiveStyle = button.DefaultStyle
	})

	button.Transform = NewTransform(0, 0)
	canvas.AssignTransform(button.Transform)
	control.AssignTransform(button.Transform)

	button.Control = control
	button.Canvas = canvas

	return button
}

func (b *button) AssignTransform(t *Transform) {
	b.Transform = t
	b.Canvas.AssignTransform(t)
	b.Control.AssignTransform(t)
}
