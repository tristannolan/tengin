package main

import (
	"fmt"
	"log"

	"github.com/tristannolan/tengin/tengin"
)

func main() {
	e, err := tengin.New()
	if err != nil {
		log.Fatalf("Failed to start tengin: %s", err)
	}
	defer e.Quit()

	w, h := e.ScreenSize()

	g := NewGame(w, h)

	// Button
	buttons := []tengin.Widget{}
	{
		title := "First Button"
		x := 10
		y := 10
		defStyle := tengin.NewStyle().NewBg(200, 200, 200).NewFg(20, 20, 20)
		hoverStyle := tengin.NewStyle().NewBg(255, 255, 255).NewFg(20, 20, 20)
		activeStyle := hoverStyle

		btn := tengin.NewButton(title, defStyle, 1)
		btn.AssignTransform(tengin.NewTransform(x, y))
		btn.HoverStyle = hoverStyle
		btn.ActiveStyle = activeStyle

		btn.Control().SetZ(1)
		btn.Control().SetClickAction(func() {
			tengin.ConsoleLog(fmt.Sprintf("%s clicked", title))
		})
		btn.Control().SetHoverAction(func() {
			btn.ActiveStyle.CopyValues(btn.HoverStyle)
		})
		btn.Control().SetHoverOffAction(func() {
			btn.ActiveStyle.CopyValues(btn.DefaultStyle)
		})

		buttons = append(buttons, btn)
	}
	{
		title := "Second Button"
		x := 16
		y := 11
		defStyle := tengin.NewStyle().NewBg(250, 100, 0).NewFg(25, 10, 0)
		hoverStyle := tengin.NewStyle().NewBg(255, 255, 255).NewFg(20, 20, 20)
		activeStyle := hoverStyle

		btn := tengin.NewButton(title, defStyle, 1)
		btn.AssignTransform(tengin.NewTransform(x, y))
		btn.HoverStyle = hoverStyle
		btn.ActiveStyle = activeStyle

		btn.Control().SetZ(1)
		btn.Control().SetClickAction(func() {
			tengin.ConsoleLog(fmt.Sprintf("%s clicked", title))
		})
		btn.Control().SetHoverAction(func() {
			btn.ActiveStyle.CopyValues(btn.HoverStyle)
		})
		btn.Control().SetHoverOffAction(func() {
			btn.ActiveStyle.CopyValues(btn.DefaultStyle)
		})

		buttons = append(buttons, btn)
	}

	for _, btn := range buttons {
		g.scene.AppendCanvas(btn.Canvas())
		g.scene.AppendControl(btn.Control())
	}
	g.scene.AppendWidget(buttons...)

	if err := e.Run(g); err != nil {
		log.Fatalf("Runtime error: %s", err)
	}
}

func (g Game) Update(ctx tengin.Context) {
	switch ctx.Key().Value() {
	case "Escape":
		ctx.Quit()
	}

	mouse := ctx.MouseKey()
	lastMouse := ctx.LastMouseKey()
	mx, my := lastMouse.Position()

	ctrl := g.scene.HitTest(mx, my)
	if ctrl != nil {
		if mouse.Key() == tengin.MouseLeft {
			ctrl.Click()
		}
	}
}

func (g Game) Draw(ctx tengin.Context) {
	ctx.SubmitScene(g.scene)
}

type Game struct {
	scene  *tengin.Scene
	canvas *tengin.Canvas
}

func NewGame(screenWidth, screenHeight int) *Game {
	scene := tengin.NewScene(screenWidth, screenHeight)
	style := tengin.NewStyle().NewBg(12, 12, 12).NewFg(240, 240, 240)
	scene.SetDefaultStyle(style)

	scene.AppendCanvas(tengin.Text("Tengin - Control"))

	g := &Game{
		scene: scene,
	}

	return g
}
