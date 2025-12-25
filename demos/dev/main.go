package main

import (
	"log"

	"github.com/tristannolan/tengin/tengin"
)

type Game struct {
	scene *tengin.Scene
}

func newGame() Game {
	return Game{}
}

func (g Game) Update(ctx tengin.Context) {
	switch ctx.Key().Value() {
	case "Escape":
		ctx.Quit()
	}
}

func (g Game) Draw(ctx tengin.Context) {
	g.scene.AppendCanvas(
		tengin.Text(0, 0, "Tengin - Dev"),
		tengin.Text(10, 1, "abcdefghijklmnopqrstuv"),
	)

	ctx.SubmitScene(g.scene)
}

func main() {
	e, err := tengin.New()
	if err != nil {
		log.Fatalf("Failed to start tengin: %s", err)
	}
	defer e.Quit()

	g := newGame()
	g.scene = tengin.NewScene()

	if err := e.Run(g); err != nil {
		log.Fatalf("Runtime error: %s", err)
	}
}
