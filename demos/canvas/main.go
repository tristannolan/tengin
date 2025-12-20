package main

import (
	"log"

	"github.com/gdamore/tcell/v3"
	"github.com/tristannolan/tengin/tengin"
)

type Game struct {
	canvases []*tengin.Canvas
}

func NewGame() Game {
	return Game{
		canvases: []*tengin.Canvas{},
	}
}

func (g Game) Update(ctx tengin.Context) {
	if ctx.Key() == tcell.KeyEscape || ctx.Key() == tcell.KeyCtrlC {
		ctx.Quit()
	}

	if ctx.Tick()%8 != 0 {
		return
	}

	for i := range g.canvases {
		g.canvases[i].X += 1
	}
}

func (g Game) Draw(ctx tengin.Context) {
	ctx.PutStr(0, 0, "Tengin - Canvas")

	scene := ctx.NewScene()

	for i := range g.canvases {
		scene.AppendCanvas(g.canvases[i])
	}

	ctx.SubmitScene(scene)
}

func main() {
	e, err := tengin.New()
	if err != nil {
		log.Fatalf("Failed to start tengin: %s", err)
	}

	g := NewGame()

	canvas := tengin.NewCanvas(10, 10)
	canvas.X = 10
	canvas.Y = 10
	canvas.Z = 1
	canvas.LoopTiles(func(x, y int) {
		tile := tengin.NewTile("F")
		tile.Fg = tengin.NewColor(100, 100, 100)
		canvas.SetTile(x, y, &tile)
	})

	g.canvases = append(g.canvases, &canvas)

	if err := e.Run(g); err != nil {
		log.Fatalf("Runtime error: %s", err)
	}
}
