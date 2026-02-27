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

	d := driver{}

	if err := e.Run(d); err != nil {
		log.Fatalf("runtime error: %s", err)
	}
}

type driver struct{}

func (g driver) Update(ctx *tengin.Context) {
}

func (g driver) Draw(ctx *tengin.Context) {
}
