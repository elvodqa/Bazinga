package main

import (
	"bazinga/core"
	"bazinga/graphics"
	"fmt"
	"runtime"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func init() {
	runtime.LockOSThread()
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()
	if err := ttf.Init(); err != nil {
		panic(err)
	}
}

func main() {

	window, err := sdl.CreateWindow("Bazinga", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 1280, 736, sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()
	renderer.Present()

	textRenderer, err := graphics.NewTextRenderer(renderer, "assets/fonts/Roboto-Regular.ttf", 24, sdl.Color{255, 255, 255, 255})
	if err != nil {
		panic(err)
	}

	imageRenderer, err := graphics.NewImageRenderer(renderer, "assets/images/peppy.jpeg")

	gameTime := core.NewGameTime()

	i := 0
	running := true
	for running {

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
			}
		}

		renderer.SetDrawColor(0, 0, 0, 255)

		textRenderer.SetSize(15)
		textRenderer.Color = sdl.Color{0, 255, 0, 255}
		textRenderer.Render(fmt.Sprintf("FPS: %f", gameTime.GetFPS()), 0, 0)

		textRenderer.SetSize(24)
		textRenderer.Color = sdl.Color{255, 0, 0, 255}
		textRenderer.Render(fmt.Sprintf("%f", gameTime.DeltaTime), int32(i), 48)
		i += 1

		imageRenderer.RenderWithSize(400, 400, 100, 200)

		renderer.Present()
		renderer.Clear()
		gameTime.Update()
	}
}
