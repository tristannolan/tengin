package main

import (
	"log"
	"strconv"

	"github.com/tristannolan/tengin/tengin"
)

var exampleHeight = 2

type Game struct {
	examples []*tengin.Canvas
}

func newGame() *Game {
	return &Game{
		examples: []*tengin.Canvas{},
	}
}

func (g *Game) Update(ctx tengin.Context) {
	if ctx.KeyIsSpecial(tengin.KeyEscape) {
		ctx.Quit()
	}

	g.newExample("Current Key", tengin.Text(0, 0, ""))
	g.newExample("Rune", tengin.Text(0, 0, string(ctx.KeyRuneValue())))
	g.newExample("Special", tengin.Text(0, 0, strconv.Itoa(int(ctx.KeySpecialValue()))))

	g.newExample("Last Key", tengin.Text(0, 0, ""))
	g.newExample("Rune", tengin.Text(0, 0, string(ctx.KeyRuneValue())))
	g.newExample("Special", tengin.Text(0, 0, strconv.Itoa(int(ctx.KeySpecialValue()))))
	g.newExample("Key Empty", tengin.Text(0, 0, string(ctx.KeyRuneValue())))
}

func (g *Game) Draw(ctx tengin.Context) {
	scene := ctx.NewScene()

	title := tengin.Text(0, 0, "Tengin - Input")
	scene.AppendCanvas(&title)

	for i := range g.examples {
		scene.AppendCanvas(g.examples[i])
	}
	g.examples = []*tengin.Canvas{}

	ctx.SubmitScene(scene)
}

func (g *Game) newExample(name string, c tengin.Canvas) {
	text := tengin.Text(0, exampleHeight, name)
	c.X = 15
	c.Y = exampleHeight
	exampleHeight += c.Height() + 1
	g.examples = append(g.examples, &text, &c)
}

func main() {
	e, err := tengin.New()
	if err != nil {
		log.Fatalf("Failed to start tengin: %s", err)
	}
	defer e.Quit()

	// Slow tick rate so we can comprehend the events
	e.SetTickRate(2)

	g := newGame()

	if err := e.Run(g); err != nil {
		log.Fatalf("Runtime error: %s", err)
	}
}
