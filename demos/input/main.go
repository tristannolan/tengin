package main

import (
	"log"
	"strconv"

	"github.com/tristannolan/tengin/tengin"
)

var (
	tickRate   = 10
	col        = 13
	textYCount = -2
)

type Game struct{}

func newGame() *Game {
	return &Game{}
}

func (g *Game) Update(ctx tengin.Context) {
	if ctx.Key().SpecialValue() == tengin.KeyEscape {
		ctx.Quit()
	}
}

func heading(t string) *tengin.Canvas {
	textYCount += 2
	return tengin.Text(0, textYCount, t)
}

func row(name, value string) *tengin.Canvas {
	textYCount++
	c := tengin.NewCanvas(0, 0, 100, 100)
	c.AppendChild(
		tengin.Text(0, textYCount, name),
		tengin.Text(col, textYCount, value),
	)
	return c
}

func (g *Game) Draw(ctx tengin.Context) {
	scene := ctx.NewScene()

	title := tengin.Text(0, 0, "Tengin - Input")
	scene.AppendCanvas(title)

	key := ctx.Key()
	lastKey := ctx.LastKey()

	mouse := ctx.MouseKey()
	mouseX, mouseY := mouse.Position()

	lastMouse := ctx.LastMouseKey()
	lastMouseX, lastMouseY := lastMouse.Position()

	textYCount = -2
	info := tengin.NewCanvas(0, 2, 100, 100)
	info.AppendChild(
		heading("Current Key"),
		row("Value", key.Value()),
		row("Special", strconv.Itoa(int(key.SpecialValue()))),
		row("Empty", strconv.FormatBool(key.IsEmpty())),

		heading("Last Key"),
		row("Value", lastKey.Value()),
		row("Special", strconv.Itoa(int(lastKey.SpecialValue()))),
		row("Empty", strconv.FormatBool(lastKey.IsEmpty())),

		heading("Mouse"),
		row("X", strconv.Itoa(mouseX)),
		row("Y", strconv.Itoa(mouseY)),
		row("Key Name", mouse.KeyName()),
		row("Key Code", strconv.Itoa(int(mouse.Key()))),
		row("Wheel Name", mouse.WheelName()),

		heading("Last Mouse"),
		row("X", strconv.Itoa(lastMouseX)),
		row("Y", strconv.Itoa(lastMouseY)),
		row("Key Name", lastMouse.KeyName()),
		row("Key Code", strconv.Itoa(int(lastMouse.Key()))),
		row("Wheel Name", lastMouse.WheelName()),

		heading("Screen"),
		row("Resizing", strconv.FormatBool(ctx.ScreenResizing())),
		row("Focused", strconv.FormatBool(ctx.ScreenFocused())),
	)
	scene.AppendCanvas(info)

	ctx.SubmitScene(scene)
}

func main() {
	e, err := tengin.New()
	if err != nil {
		log.Fatalf("Failed to start tengin: %s", err)
	}
	defer e.Quit()

	// Events happen fast so let's slow the tick rate
	e.SetTickRate(tickRate)

	g := newGame()

	if err := e.Run(g); err != nil {
		log.Fatalf("Runtime error: %s", err)
	}
}
