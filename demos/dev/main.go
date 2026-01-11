package main

import (
	"log"
	"strconv"

	"github.com/tristannolan/tengin/tengin"
)

type Game struct {
	scene   *tengin.Scene
	text    *tengin.Canvas
	wrapper *tengin.Canvas
}

func newGame() Game {
	return Game{}
}

func main() {
	e, err := tengin.New()
	if err != nil {
		log.Fatalf("Failed to start tengin: %s", err)
	}
	defer e.Quit()

	e.SetTickRate(1)
	e.SetFrameRate(1)

	g := newGame()

	w, h := e.ScreenSize()
	g.scene = tengin.NewScene(w, h)
	g.wrapper = tengin.Box(10, 10, tengin.NewColor(100, 100, 100))
	g.text = tengin.Text("T")

	g.wrapper.AppendChild(g.text)
	g.scene.AppendCanvas(g.wrapper)

	if err := e.Run(g); err != nil {
		log.Fatalf("Runtime error: %s", err)
	}
}

func (g Game) Update(ctx tengin.Context) {
}

var count = 0

func (g Game) Draw(ctx tengin.Context) {
	if ctx.Tick()%4 == 0 {
		count++
		g.text.SetTile(0, 0, tengin.NewTile(strconv.Itoa(count), tengin.NewStyle()))
		if count >= 9 {
			count = 0
		}
	}
	ctx.SubmitScene(g.scene)
}
