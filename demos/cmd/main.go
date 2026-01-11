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

	g := newGame(w, h)

	if err := e.Run(g); err != nil {
		log.Fatalf("Runtime error: %s", err)
	}
}

type Game struct {
	scene  *tengin.Scene
	canvas *tengin.Canvas
}

func newGame(screenWidth, screenHeight int) *Game {
	scene := tengin.NewScene(screenWidth, screenHeight)
	style := tengin.NewStyle().NewBg(12, 12, 12).NewFg(240, 240, 240)
	scene.SetDefaultStyle(style)

	scene.AppendCanvas(tengin.Text("Tengin - Command"))

	g := &Game{
		scene: scene,
	}

	return g
}

func (g Game) Update(ctx tengin.Context) {
}

func (g Game) Draw(ctx tengin.Context) {
	ctx.SubmitScene(g.scene)
}
