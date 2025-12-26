package main

import (
	"log"
	"strconv"

	"github.com/tristannolan/tengin/tengin"
)

var (
	updateCount = 0
	drawCount   = 0
)

type Game struct {
	scene *tengin.Scene
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

	g := newGame()
	g.scene = tengin.NewScene()

	if err := e.Run(g); err != nil {
		log.Fatalf("Runtime error: %s", err)
	}
}

func (g Game) Update(ctx tengin.Context) {
	switch ctx.Key().Value() {
	case "Escape":
		ctx.Quit()
	}

	heavy()

	updateCount++
}

func (g Game) Draw(ctx tengin.Context) {
	drawCount++

	g.scene.AppendCanvas(
		tengin.Text(0, 0, "Tengin - Dev"),
		tengin.Text(0, 1, "Update Count"),
		tengin.Text(20, 1, strconv.Itoa(updateCount)),
		tengin.Text(0, 2, "Draw Count"),
		tengin.Text(20, 2, strconv.Itoa(drawCount)),
	)

	drawCount = 0
	updateCount = 0

	ctx.SubmitScene(g.scene)

	tengin.DebugLog("Load", load)
}

var load = 6000

func heavy() {
	for range load {
		for range load {
		}
	}
	load += 1
}
