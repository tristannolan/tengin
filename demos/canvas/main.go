package main

import (
	"log"

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
	if ctx.Key().SpecialValue() == tengin.KeyEscape {
		ctx.Quit()
	}
}

func (g *Game) Draw(ctx tengin.Context) {
	scene := ctx.NewScene()

	title := tengin.Text(0, 0, "Tengin - Canvas")
	scene.AppendCanvas(title)

	for i := range g.examples {
		scene.AppendCanvas(g.examples[i])
	}

	ctx.SubmitScene(scene)
}

func (g *Game) newExample(name string, c *tengin.Canvas) {
	text := tengin.Text(0, exampleHeight, name)
	c.X = 15
	c.Y = exampleHeight
	exampleHeight += c.Height() + 1
	g.examples = append(g.examples, text, c)
}

func main() {
	e, err := tengin.New()
	if err != nil {
		log.Fatalf("Failed to start tengin: %s", err)
	}
	defer e.Quit()

	g := newGame()

	g.newExample("Text",
		tengin.Text(0, 0, "Write something funny"),
	)
	g.newExample("Paragraph",
		tengin.Paragraph(0, 0, 40, "Paragraph will box your text in.\n\n It will split text into words, cap line width, and insert newlines where necessary."),
	)
	g.newExample("Box",
		tengin.Box(0, 0, 40, 3, tengin.NewColor(100, 150, 150)),
	)
	g.newExample("Nesting",
		newParentExample(10, 10, 2),
	)

	if err := e.Run(g); err != nil {
		log.Fatalf("Runtime error: %s", err)
	}
}

func newParentExample(x, y, z int) *tengin.Canvas {
	parent := tengin.Box(x, y, 40, 10, tengin.NewColor(100, 150, 150))
	parent.Z = z

	child1 := tengin.Box(1, 1, 20, 5, tengin.NewColor(70, 130, 130))
	child1.Z = 3

	child2 := tengin.Box(15, 2, 15, 5, tengin.NewColor(60, 120, 120))
	child2.Z = 2

	parent.AppendChild(child1, child2)
	return parent
}
