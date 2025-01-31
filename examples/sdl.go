package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	winWidth  = 800
	winHeight = 600
)

func main() {
	// Initialize SDL
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("Error initializing SDL:", err)
		return
	}
	defer sdl.Quit()

	// Create window
	window, err := sdl.CreateWindow("SDL2 in Go!", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println("Could not create window:", err)
		return
	}
	defer window.Destroy()

	// Create renderer
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Could not create renderer:", err)
		return
	}
	defer renderer.Destroy()

	// Main loop flag
	running := true

	// Event loop
	for running {
		// Handle events
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event := event.(type) {
			case *sdl.QuitEvent: // Quit button clicked
				running = false
			case *sdl.KeyboardEvent: // Handle keyboard
				if event.Type == sdl.KEYDOWN {
					if event.Keysym.Sym == sdl.K_ESCAPE { // ESC key to exit
						running = false
					}
				}
			}
		}

		// Set background color (light blue)
		renderer.SetDrawColor(173, 216, 230, 255)
		renderer.Clear()

		// Draw a red rectangle
		renderer.SetDrawColor(255, 0, 0, 255)
		rect := sdl.Rect{X: 300, Y: 200, W: 200, H: 150}
		renderer.FillRect(&rect)

		// Show the result
		renderer.Present()

		// Limit frame rate to ~60 FPS
		sdl.Delay(16)
	}
}
