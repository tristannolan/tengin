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

	g := newGame()

	if err := e.Run(g); err != nil {
		log.Fatalf("Runtime error: %s", err)
	}
}

type Game struct {
	scene *tengin.Scene
}

func newGame() *Game {
	scene := tengin.NewScene()
	scene.SetDefaultStyle(tengin.NewStyle().NewBg(10, 10, 10))
	return &Game{
		scene: scene,
	}
}

func (g Game) Update(ctx tengin.Context) {
	switch ctx.Key().Value() {
	case "p":
	// Plant crop
	case "h":
	// Harvest crop
	case "Escape":
		ctx.Quit()
	}
}

func (g Game) Draw(ctx tengin.Context) {
	g.scene.AppendCanvas(tengin.Text(0, 0, "Tengin - Basic Instance"))
	ctx.SubmitScene(g.scene)
}
