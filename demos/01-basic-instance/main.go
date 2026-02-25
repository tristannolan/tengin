package main

import (
	"log"

	"github.com/tristannolan/tengin/tengin"
)

func main() {
	e, err := tengin.New()
	if err != nil {
		log.Fatalf("failed to start tengin: %s", err)
	}
	defer e.Quit()

	g := &Game{}

	if err := e.Run(g); err != nil {
		log.Fatalf("runtime error: %s", err)
	}
}

type Game struct{}

func (g Game) Update(ctx tengin.Context) {
}

func (g Game) Draw(ctx tengin.Context) {
}
