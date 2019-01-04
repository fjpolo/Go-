package main

/* Import */
import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

/* Homework */
// Frame independance
// Score
// Game over state
// 2 player vs PC
// Handling resizing windows
// Mouse/Joystick
// Imperfect AI

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

/* Position */
type pos struct {
	x, y float32
}

/* Ball */
type ball struct {
	pos            // Position
	radius int     // radius
	xv     float32 // x velocity
	yv     float32 // y velocity
	colour colour
}

/* Paddle */
type paddle struct {
	pos           // Position
	w      int    // height
	h      int    // width
	colour colour // Colour
}

/*****************************
 *
 * Methods
 *
 *****************************/

/* Paddle draw */
func (paddle *paddle) draw(pixels []byte) {
	/* Starting point */
	startX := int(paddle.x) - paddle.w/2
	startY := int(paddle.y) - paddle.h/2

	/* Draw the rectangle for the paddle */
	for y := 0; y < paddle.h; y++ {
		for x := 0; x < paddle.w; x++ {
			setPixel(startX+x, startY+y, paddle.colour, pixels)
		}
	}
}

/* Paddle update */
func (paddle *paddle) update(keyState []uint8) {
	/* UP arrow */
	if keyState[sdl.SCANCODE_UP] != 0 {
		paddle.y -= 5
	}
	/* DOWN arrow */
	if keyState[sdl.SCANCODE_DOWN] != 0 {
		paddle.y += 5
	}
}

/* Paddle bot */
func (paddle *paddle) aiUpdate(ball *ball) {
	/* Unbeatable */
	paddle.y = ball.y
}

/* Ball draw */
func (ball *ball) draw(pixels []byte) {

	/* Iterate over a square and draw if within radius */
	for y := -ball.radius; y < ball.radius; y++ {
		for x := -ball.radius; x < ball.radius; x++ {
			/* Check radius boundaries */
			if x*x+y*y < ball.radius*ball.radius {
				/*  */
				setPixel(int(ball.x)+x, int(ball.y)+y, ball.colour, pixels)
			}
		}
	}
}

/* Ball update */
func (ball *ball) update(leftPaddle *paddle, rightPaddle *paddle) {
	/**/
	ball.x += ball.xv
	ball.y += ball.yv

	/* Screen boundaries - Top and Bottom */
	if (int(ball.y)-ball.radius < 0) || (int(ball.y)+ball.radius > winHeight) {
		/* Reverse velocity */
		ball.yv = -ball.yv
	}
	/* Screen boundaries - Left and Right */
	if (int(ball.x)-ball.radius < 0) || (int(ball.x)+ball.radius > winWidth) {
		/* Center ball */
		ball.x = 300
		ball.y = 300
	}

	/* Hit a paddle left */
	if int(ball.x)-ball.radius < int(leftPaddle.x)+leftPaddle.w/2 {
		/**/
		if int(ball.y) > int(leftPaddle.y)-leftPaddle.h/2 &&
			int(ball.y) < int(leftPaddle.y)+leftPaddle.h/2 {
			/* Reverse velocity */
			ball.xv = -ball.xv
		}
	}
	/* Hit a paddle right */
	if int(ball.x)+ball.radius > int(rightPaddle.x)-rightPaddle.w/2 {
		/**/
		if int(ball.y) > int(rightPaddle.y)-rightPaddle.h/2 &&
			int(ball.y) < int(rightPaddle.y)+rightPaddle.h/2 {
			/* Reverse velocity */
			ball.xv = -ball.xv
		}
	}
}

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

/* Clear screen */
func clear(pixels []byte) {
	/**/
	for i := range pixels {
		pixels[i] = 0
	}
}

/*****************************
 *
 *          MAIN
 *
 *****************************/
func main() {

	/* Initialise everything */
	err := sdl.Init(sdl.INIT_EVERYTHING)
	/* Check error */
	if err != nil {
		fmt.Println(err)
		return
	}

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

	/* Player paddles*/
	player1 := paddle{pos{50, 300}, 20, 100, colour{255, 255, 255}}
	player2 := paddle{pos{float32(winWidth - 50), 300}, 20, 100, colour{255, 255, 255}}

	/* Ball */
	ball := ball{pos{400, 300}, 10, 3, 3, colour{255, 255, 255}}

	/* Keyboard state */
	keyState := sdl.GetKeyboardState()

	/**
	 *
	 *    Game loop
	 *
	 **/
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			/**/
			switch event.(type) {
			/* Quit */
			case *sdl.QuitEvent:
				return
				/* Keyboard event */
			}
		}
		/* Clear screen */
		clear(pixels)

		/* Updates */
		player1.update(keyState)
		player2.aiUpdate(&ball)
		ball.update(&player1, &player2)

		/* Draw paddles*/
		player1.draw(pixels)
		player2.draw(pixels)
		/* Draw ball */
		ball.draw(pixels)

		/*Update texture*/
		texture.Update(nil, pixels, winWidth*4)
		renderer.Copy(texture, nil, nil)
		renderer.Present()

		/* Delay */
		sdl.Delay(16)
	}
}
