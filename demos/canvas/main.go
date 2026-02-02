package main

import (
	"log"

	"github.com/tristannolan/tengin/tengin"
)

var exampleHeight = 2

func main() {
	e, err := tengin.New()
	if err != nil {
		log.Fatalf("Failed to start tengin: %s", err)
	}
	defer e.Quit()

	e.SetTickRate(60)
	e.SetFrameRate(60)

	w, h := e.ScreenSize()
	g := newGame(w, h)

	title := tengin.Text("Tengin - Canvas")

	g.newExample("Text",
		tengin.Text("Write something funny"),
	)
	g.newExample("Paragraph",
		tengin.Paragraph(40, "Box text in by splitting words and capping line width.\n\n Now with newlines."),
	)
	g.newExample("Box",
		tengin.Box(40, 3, tengin.NewColor(100, 150, 150)),
	)
	g.newExample("Nesting",
		newParentExample(10, 10, 2),
	)

	g.scene.AppendCanvas(title)
	for i := range g.examples {
		g.scene.AppendCanvas(g.examples[i])
	}

	if err := e.Run(g); err != nil {
		log.Fatalf("Runtime error: %s", err)
	}
}

type Game struct {
	examples []*tengin.Canvas
	scene    *tengin.Scene
}

func newGame(screenWidth, screenHeight int) *Game {
	scene := tengin.NewScene(screenWidth, screenHeight)
	style := tengin.NewStyle().NewBg(12, 12, 12).NewFg(240, 240, 240)
	scene.SetDefaultStyle(style)

	return &Game{
		examples: []*tengin.Canvas{},
		scene:    scene,
	}
}

func (g *Game) Update(ctx tengin.Context) {
}

func (g *Game) Draw(ctx tengin.Context) {
	ctx.SubmitScene(g.scene)
}

func (g *Game) newExample(name string, c *tengin.Canvas) {
	text := tengin.Text(name)
	text.Position(0, exampleHeight)
	// c.X = 15
	// c.Y = exampleHeight
	c.Transform(15, exampleHeight)
	exampleHeight += c.Height + 1
	g.examples = append(g.examples, text, c)
}

func newParentExample(tx, ty, z int) *tengin.Canvas {
	parent := tengin.Box(40, 10, tengin.NewColor(100, 150, 150))
	parent.Transform(tx, ty)
	parent.SetZ(z)

	child1 := tengin.Box(20, 5, tengin.NewColor(70, 130, 130))
	child1.Position(1, 1)
	child1.SetZ(3)

	child2 := tengin.Box(15, 5, tengin.NewColor(60, 120, 120))
	child2.Position(15, 2)
	child2.SetZ(2)

	parent.AppendChild(child1, child2)
	return parent
}
