package main

import (
	"log"

	"github.com/tristannolan/tengin/tengin"
)

type driver struct{}

func main() {
	e, err := tengin.New()
	if err != nil {
		log.Fatalf("failed to start tengin: %s", err)
	}
	defer e.Stop()

	d := driver{}

	if err := e.Run(d); err != nil {
		log.Fatalf("runtime error: %s", err)
	}
}

func (g driver) Update(ctx *tengin.Context) {
}

func (g driver) Draw(ctx *tengin.Context) {
}
