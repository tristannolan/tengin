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

	e.SetTickRate(60)
	e.SetFrameRate(60)

	w, h := e.ScreenSize()

	g := newGame(w, h)
	g.scene.AppendCanvas(tengin.Text(0, 0, "Tengin - Basic Instance"))

	if err := e.Run(g); err != nil {
		log.Fatalf("Runtime error: %s", err)
	}
}

type Game struct {
	scene *tengin.Scene
}

func newGame(screenWidth, screenHeight int) *Game {
	return &Game{
		scene: tengin.NewScene(screenWidth, screenHeight),
	}
}

func (g Game) Update(ctx tengin.Context) {
	switch ctx.Key().Value() {
	case "Escape":
		ctx.Quit()
	}
}

func (g Game) Draw(ctx tengin.Context) {
	ctx.SubmitScene(g.scene)
}
