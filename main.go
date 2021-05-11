package main

import (
	"sdl-test/config"

	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {

	conf, err := config.Load()

	if err != nil {
		panic(err)
	}

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		conf.Dimensions.Width+conf.Border*2, conf.Dimensions.Height+conf.Border*2, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	_, err = window.GetSurface()

	if err != nil {
		panic(err)
	}

	renderer, err := window.GetRenderer()

	if err != nil {
		panic(err)
	}

	//surface.FillRect(nil, 0)
	renderer.DrawRect(nil)

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
			drawBoard(renderer, conf)
			renderer.Present()

			// sdl.Delay(16)
		}
	}
}

func holeSize(conf *config.Connect4Config) (int32, int32) {
	holeWidth := (conf.Dimensions.Width - (conf.Border*2 + conf.Gap*(conf.Columns-1) + 2*conf.Padding)) / conf.Columns
	holeHeight := (conf.Dimensions.Height - (conf.Border*2 + conf.Gap*(conf.Rows-1) + 2*conf.Padding)) / conf.Rows
	return holeWidth, holeHeight
}

func drawBoard(renderer *sdl.Renderer, conf *config.Connect4Config) {
	//rect := sdl.Rect{X: conf.Border, Y: conf.Border, W: conf.Dimensions.Width, H: conf.Dimensions.Height}

	holeWidth, holeHeight := holeSize(conf)
	gfx.BoxColor(renderer, conf.Border, conf.Border, conf.Dimensions.Width, conf.Dimensions.Height, sdl.Color{
		R: 0,
		G: 0,
		B: 255,
		A: 255,
	})

	xOffset := conf.Border + conf.Padding
	yOffset := conf.Border + conf.Padding
	for rowI := int32(0); rowI < conf.Rows; rowI++ {
		for colI := int32(0); colI < conf.Columns; colI++ {
			gfx.FilledEllipseColor(renderer,
				xOffset+holeWidth/2+(conf.Gap+holeWidth)*colI,
				yOffset+holeHeight/2+(conf.Gap+holeHeight)*rowI,
				holeWidth/2, holeHeight/2, sdl.Color{R: 255, G: 255, B: 255, A: 255})
		}
	}

}
