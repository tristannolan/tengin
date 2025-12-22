package main

import (
	"log"

	"github.com/gdamore/tcell/v3"
	"github.com/tristannolan/tengin/tengin"
)

type Game struct {
	title *tengin.Canvas
}

func newGame() Game {
	return Game{}
}

func (g Game) Update(ctx tengin.Context) {
	if ctx.Key() == tcell.KeyEscape || ctx.Key() == tcell.KeyCtrlC {
		ctx.Quit()
	}
}

func (g Game) Draw(ctx tengin.Context) {
	scn := ctx.NewScene()
	scn.AppendCanvas(g.title)
	ctx.SubmitScene(scn)
}

func main() {
	e, err := tengin.New()
	if err != nil {
		log.Fatalf("Failed to start tengin: %s", err)
	}
	defer e.Quit()

	g := newGame()
	text := tengin.Text(0, 0, "Tengin - Basic Instance")
	g.title = &text

	if err := e.Run(g); err != nil {
		log.Fatalf("Runtime error: %s", err)
	}
}
