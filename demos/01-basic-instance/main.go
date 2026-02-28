package main

import (
	"log"

	"github.com/tristannolan/tengin/tengin"
)

type driver struct{}

func main() {
	// A new instance is created with the default config.
	e, err := tengin.New()
	if err != nil {
		log.Fatalf("failed to start tengin: %s", err)
	}
	defer e.Stop()

	// We'll add our own custom config.
	e.LoadConfig(tengin.Config{
		Title:     "Tengin - Basic Instance",
		TickRate:  60,
		FrameRate: 60,
	})

	// The driver receives context through update() and draw().
	// In reality this would be the main game/program struct.
	d := driver{}

	// Running the engine will complete initialisation and start the loop.
	if err := e.Run(d); err != nil {
		log.Fatalf("runtime error: %s", err)
	}
}

func (g driver) Update(ctx *tengin.Context) {
}

func (g driver) Draw(ctx *tengin.Context) {
}
