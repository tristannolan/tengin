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

	if ctx.Tick()%120 != 0 {
		return
	}

	for i := range g.canvases {
		g.canvases[i].X += 1
	}
}

func (g Game) Draw(ctx tengin.Context) {
	scene := ctx.NewScene()

	for i := range g.canvases {
		scene.AppendCanvas(g.canvases[i])
	}

	ctx.SubmitScene(scene)
}

func NewCanvasBox(x, y, width, height int, clr tengin.Color) tengin.Canvas {
	canvas := tengin.NewCanvas(x, y, width, height)
	for y := range canvas.Tiles {
		for x := range canvas.Tiles[y] {
			tile := tengin.NewTile("O")
			tile.Fg = clr
			canvas.SetTile(x, y, &tile)
		}
	}
	return canvas
}

func NewParentExample(x, y, z, clr int) tengin.Canvas {
	parent := NewCanvasBox(x, y, 10, 10, tengin.NewColor(225, 80, int32(clr)))
	parent.Z = z
	child := NewCanvasBox(1, 1, 5, 5, tengin.NewColor(80, 225, int32(clr)))
	child.Z = z + 3
	parent.AppendChild(&child)
	return parent
}

func main() {
	e, err := tengin.New()
	if err != nil {
		log.Fatalf("Failed to start tengin: %s", err)
	}

	g := NewGame()

	parent1 := NewParentExample(10, 10, 2, 0)
	parent2 := NewParentExample(15, 15, 1, 255)

	g.canvases = append(g.canvases, &parent1, &parent2)

	if err := e.Run(g); err != nil {
		log.Fatalf("Runtime error: %s", err)
	}
}
