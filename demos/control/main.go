package main

import (
	"log"

	"github.com/tristannolan/tengin/tengin"
)

func main() {
	e, err := tengin.New()
	if err != nil {
		log.Fatalf("Failed to start tengin: %s", err)
	}
	defer e.Quit()

	// e.SetTickRate(1)
	// e.SetFrameRate(1)

	w, h := e.ScreenSize()

	g := NewGame(w, h)

	// Button

	// Button Canvas
	buttonStyle1 := tengin.NewStyle().NewBg(200, 200, 200).NewFg(20, 20, 20)
	button1 := tengin.NewButton("First Button", buttonStyle1, 1)
	button1.HoverStyle = tengin.NewStyle().NewBg(255, 255, 255).NewFg(20, 20, 20)
	button1.AssignTransform(tengin.NewTransform(10, 10))
	button1.Control.SetZ(1)
	button1.Control.SetClickAction(func() {
		tengin.ConsoleLog("First button clicked")
	})
	button1.Control.SetHoverAction(func() {
		button1.ActiveStyle.CopyValues(button1.HoverStyle)
	})
	button1.Control.SetHoverOffAction(func() {
		button1.ActiveStyle.CopyValues(button1.DefaultStyle)
	})

	buttonStyle2 := tengin.NewStyle().NewBg(250, 100, 0).NewFg(25, 10, 0)
	button2 := tengin.NewButton("Second Button", buttonStyle2, 1)
	button2.HoverStyle = tengin.NewStyle().NewBg(255, 255, 255).NewFg(20, 20, 20)
	button2.AssignTransform(tengin.NewTransform(16, 11))
	button2.Control.SetZ(2)
	button2.Control.SetClickAction(func() {
		tengin.ConsoleLog("Second button clicked")
		button2.Canvas.Transform(1, 1)
	})
	button2.Control.SetHoverAction(func() {
		button2.ActiveStyle.CopyValues(button2.HoverStyle)
	})
	button2.Control.SetHoverOffAction(func() {
		button2.ActiveStyle.CopyValues(button2.DefaultStyle)
	})

	g.scene.AppendCanvas(button1.Canvas, button2.Canvas)
	g.scene.AppendControl(button1.Control, button2.Control)

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
