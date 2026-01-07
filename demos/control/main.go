package main

import (
	"log"
	"strings"

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

	// Button Canvas
	canvas1, control1 := newButton(10, 10, "First Button", 200, 200, 200)
	control1.SetZ(2)
	canvas2, control2 := newButton(16, 11, "The Second", 250, 100, 0)
	control2.SetZ(1)
	g.scene.AppendCanvas(canvas1, canvas2)
	g.scene.AppendControl(control1, control2)

	if err := e.Run(g); err != nil {
		log.Fatalf("Runtime error: %s", err)
	}
}

func newButton(x, y int, msg string, r, g, b int32) (*tengin.Canvas, *tengin.Control) {
	chars := strings.Split(msg, "")
	padding := 1
	width := len(msg) + padding*2
	height := 1 + padding*2

	canvas := tengin.Box(x, y, width, height, tengin.NewColor(r, g, b))
	style := tengin.NewStyle().NewFg(r/10, g/10, b/10)
	for y := range canvas.Tiles {
		for _, tile := range canvas.Tiles[y] {
			tile.Char = " "
		}
	}
	for i, char := range chars {
		tile := tengin.NewTile(char, style)
		canvas.SetTile(i+padding, padding, tile)
	}
	canvas.SetAlwaysCache(true)

	// Button Control
	ctrl := tengin.NewControl(x, y, width, height, func() {
		tengin.ConsoleLog("Button press:" + msg)
	})

	return canvas, ctrl
}

type Game struct {
	scene  *tengin.Scene
	canvas *tengin.Canvas
}

func NewGame(screenWidth, screenHeight int) *Game {
	scene := tengin.NewScene(screenWidth, screenHeight)
	style := tengin.NewStyle().NewBg(12, 12, 12).NewFg(240, 240, 240)
	scene.SetDefaultStyle(style)

	scene.AppendCanvas(tengin.Text(0, 0, "Tengin - Control"))

	g := &Game{
		scene: scene,
	}

	return g
}

func (g Game) Update(ctx tengin.Context) {
	switch ctx.Key().Value() {
	case "Escape":
		ctx.Quit()
	}

	mx, my := ctx.MouseKey().Position()

	ctrl := g.scene.HitTest(mx, my)
	if ctrl != nil && ctx.MouseKey().Key() == tengin.MouseLeft {
		ctrl.Action()
	}
}

func (g Game) Draw(ctx tengin.Context) {
	ctx.SubmitScene(g.scene)
}
