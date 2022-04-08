package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
	"github.com/go-vgo/robotgo"
)

func main() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press ESC to quit")
	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		MoveMouse(key)
		if key == keyboard.KeyEsc {
			break
		}
	}
}

func MoveMouse(key keyboard.Key) {
	speed := 30
	xAxisTransform := 0
	yAxisTransform := 0

	if key == keyboard.KeyArrowUp {
		yAxisTransform += (speed * -1)
	} else if key == keyboard.KeyArrowLeft {
		xAxisTransform += (speed * -1)
	} else if key == keyboard.KeyArrowDown {
		yAxisTransform += speed
	} else if key == keyboard.KeyArrowRight {
		xAxisTransform += speed
	}

	robotgo.MoveRelative(xAxisTransform, yAxisTransform)
}
