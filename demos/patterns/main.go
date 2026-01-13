package main

import (
	"log"
	"path/filepath"

	"github.com/tristannolan/tengin/tengin"
)

func main() {
	e, err := tengin.New()
	if err != nil {
		log.Fatalf("Failed to start tengin: %s", err)
	}
	defer e.Quit()

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

	scene.AppendCanvas(tengin.Text("Tengin - Patterns"))

	path := filepath.Join("patterns", "test-pattern")
	defStyle := tengin.NewStyle().NewFg(240, 240, 240)
	styles := map[string]*tengin.Style{
		"r": tengin.NewStyle().NewFg(180, 40, 60),
		"b": tengin.NewStyle().NewFg(60, 40, 180),
		"t": tengin.NewStyle(),
	}
	pattern, err := tengin.LoadPattern(path, defStyle, styles)
	if err != nil {
		panic(err)
	}

	pattern.Canvas.Transform(10, 10)
	scene.AppendCanvas(pattern.Canvas)
	tengin.ConsoleLogF("Len args: %d", len(pattern.Args))
	tengin.ConsoleLogF("Len arg phrases: %d", len(pattern.ArgsPhrases))

	for name, arg := range pattern.Args {
		tengin.ConsoleLogF("Arg (%s) has length %d", name, len(arg))
	}

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
