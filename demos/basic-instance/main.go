package main

import (
	"log"

	"github.com/gdamore/tcell/v3"
	"github.com/tristannolan/tengin/tengin"
)

type Game struct{}

func NewGame() Game {
	return Game{}
}

func (g Game) Update(ctx tengin.Context) {
	if ctx.Key() == tcell.KeyEscape || ctx.Key() == tcell.KeyCtrlC {
		ctx.Quit()
	}
}

func (g Game) Draw(ctx tengin.Context) {
	ctx.Clear()

	ctx.PutStr(0, 0, ctx.Str())
	ctx.Show()
}

func main() {
	e, err := tengin.New()
	if err != nil {
		log.Fatalf("Failed to start tengin: %s", err)
	}

	g := NewGame()

	if err := e.Run(g); err != nil {
		log.Fatalf("Runtime error: %s", err)
	}
}
