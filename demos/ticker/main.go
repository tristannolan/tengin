package main

import (
	"fmt"
	"time"
)

func main() {
	running := true

	updateChan := make(chan struct{})
	drawChan := make(chan struct{})

	// User process
	go func() {
		for running {
			select {
			case <-updateChan:
				fmt.Println("Update")
			case <-drawChan:
				fmt.Println("Draw")
			}
		}
	}()

	// Update
	go func() {
		for running {
			time.Sleep(10 * time.Millisecond)
			updateChan <- struct{}{}
		}
	}()

	// Draw
	go func() {
		for running {
			time.Sleep(50 * time.Millisecond)
			drawChan <- struct{}{}
		}
	}()

	time.Sleep(1 * time.Second)
	running = false
}
