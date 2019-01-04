package main

/* Import */
import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

/*****************************
 *
 * Global variables
 *
 *****************************/
const winWidth int = 800
const winHeight int = 600

/*****************************
 *
 * Structures
 *
 *****************************/
/* Colour */
type colour struct {
	r, g, b byte
}

/*****************************
 *
 * Methods
 *
 *****************************/

/*****************************
 *
 * Functions
 *
 *****************************/
/*setPixel*/
func setPixel(x, y int, c colour, pixels []byte) {
	/**/
	index := (y*winWidth + x) * 4
	/**/
	if index < len(pixels)-4 && index >= 0 {
		pixels[index] = c.r
		pixels[index+1] = c.g
		pixels[index+2] = c.b
	}
}

/*****************************
 *
 *          MAIN
 *
 *****************************/
func main() {

	/*Create window*/
	window, err := sdl.CreateWindow("Testing SDL2", sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED, int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN)
	/* Check error*/
	if err != nil {
		fmt.Println(err)
		return
	}
	/* Defer windows close */
	defer window.Destroy()

	/*Renderer*/
	renderer, err := sdl.CreateRenderer(window, 1, sdl.RENDERER_ACCELERATED)
	/* Check error*/
	if err != nil {
		fmt.Println(err)
		return
	}
	/* Defer renederer destroy */
	defer renderer.Destroy()

	/* Texture */
	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_ARGB8888, sdl.TEXTUREACCESS_STREAMING,
		int32(winWidth), int32(winHeight))
	/* Check error*/
	if err != nil {
		fmt.Println(err)
		return
	}
	/* Defer texture destroy */
	defer texture.Destroy()

	/*Pixels*/
	pixels := make([]byte, winWidth*winHeight*4) // 4 bytes: Alpha, Blue, Green, Red

	/*Draw*/
	for y := 0; y < winHeight; y++ {
		for x := 0; x < winWidth; x++ {
			/* Set pixel */
			setPixel(x, y, colour{byte((x + 2*y) % 255), byte((y * x) % 255), 0}, pixels)
		}
	}
	/*Update texture*/
	texture.Update(nil, pixels, winWidth*4)
	renderer.Copy(texture, nil, nil)
	renderer.Present()

	/*Delay*/
	sdl.Delay(5000)
}
