package main

/* Import */
import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

/* Homework */
// Score
// Game over state
// 2 player vs PC
// Handling resizing windows
// Mouse/Joystick

/*****************************
 *
 * Definitions
 *
 *****************************/
const MIN_ELAPSED_TIME = 0.005
const MAX_SCORE = 3

/*****************************
 *
 * Global variables
 *
 *****************************/
const winWidth int = 800
const winHeight int = 600

/* Game State */
type gameState int

const (
	start gameState = iota
	play
	finished
	gameover
)

var state = start

/* Level */
type gameLevel int

const (
	level1 gameLevel = iota
	level2
	level3
	level4
	level5
)

var level gameLevel = level1

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
	radius float32 // radius
	xv     float32 // x velocity
	yv     float32 // y velocity
	colour colour
}

/* Paddle */
type paddle struct {
	pos            // Position
	w      float32 // height
	h      float32 // width
	speed  float32 // speed
	score  int     // score
	colour colour  // Colour
}

/* Font */
var nums = [][]byte{
	{ // 0
		1, 1, 1,
		1, 0, 1,
		1, 0, 1,
		1, 0, 1,
		1, 1, 1,
	},
	{ // 1
		1, 1, 0,
		0, 1, 0,
		0, 1, 0,
		0, 1, 0,
		1, 1, 1,
	},
	{ // 2
		1, 1, 1,
		0, 0, 1,
		1, 1, 1,
		1, 0, 0,
		1, 1, 1,
	},
	{ // 3
		1, 1, 1,
		0, 0, 1,
		0, 1, 1,
		0, 0, 1,
		1, 1, 1,
	},
	{ // 4
		1, 0, 1,
		1, 0, 1,
		1, 1, 1,
		0, 0, 1,
		0, 0, 1,
	},
	{ // 5
		1, 1, 1,
		1, 0, 0,
		1, 1, 1,
		0, 0, 1,
		1, 1, 1,
	},
	{ // 6
		1, 1, 1,
		1, 0, 0,
		1, 1, 1,
		1, 0, 1,
		1, 1, 1,
	},
	{ // 7
		1, 1, 1,
		0, 0, 1,
		0, 0, 1,
		0, 0, 1,
		0, 0, 1,
	},
	{ // 8
		1, 1, 1,
		1, 0, 1,
		1, 1, 1,
		1, 0, 1,
		1, 1, 1,
	},
	{ // 9
		1, 1, 1,
		1, 0, 1,
		1, 1, 1,
		0, 0, 1,
		1, 1, 1,
	},
}

/*****************************
 *
 * Methods
 *
 *****************************/

/* Paddle draw */
func (paddle *paddle) draw(pixels []byte) {
	/* Starting point */
	startX := int(paddle.x - paddle.w/2)
	startY := int(paddle.y - paddle.h/2)

	/* Draw the rectangle for the paddle */
	for y := 0; y < int(paddle.h); y++ {
		for x := 0; x < int(paddle.w); x++ {
			setPixel(startX+x, startY+y, paddle.colour, pixels)
		}
	}

	/* Draw score */
	numX := lerp(paddle.x, getCenter().x, 0.2)
	drawNumber(pos{numX, 35}, paddle.colour, 10, paddle.score, pixels)
}

/* Paddle update */
func (paddle *paddle) update(keyState []uint8, elapsedTime float32) {
	/* UP arrow */
	if keyState[sdl.SCANCODE_UP] != 0 {
		/* Boundaries */
		if (paddle.y - paddle.h/2) > 0 {
			paddle.y -= paddle.speed * elapsedTime
		}
	}
	/* DOWN arrow */
	if keyState[sdl.SCANCODE_DOWN] != 0 {
		/* Boundaries */
		if (paddle.y + paddle.h/2) < float32(winHeight) {
			paddle.y += paddle.speed * elapsedTime
		}
	}
}

/* Paddle bot */
func (paddle *paddle) aiUpdate(ball *ball, elapsedTime float32) {
	/* Unbeatable */
	//paddle.y = ball.y
	/* Beatable */
	if paddle.y > ball.y {
		/* Boundaries */
		if (paddle.y - paddle.h/2) > 0 {
			paddle.y -= paddle.speed * elapsedTime
		}
	} else if paddle.y < ball.y {
		/* Boundaries */
		if (paddle.y + paddle.h/2) < float32(winHeight) {
			paddle.y += paddle.speed * elapsedTime
		}
	}
}

/* Ball draw */
func (ball *ball) draw(pixels []byte) {

	/* Iterate over a square and draw if within radius */
	for y := -ball.radius; y < ball.radius; y++ {
		for x := -ball.radius; x < ball.radius; x++ {
			/* Check radius boundaries */
			if x*x+y*y < ball.radius*ball.radius {
				/*  */
				setPixel(int(ball.x+x), int(ball.y+y), ball.colour, pixels)
			}
		}
	}
}

/* Ball update */
func (ball *ball) update(leftPaddle *paddle, rightPaddle *paddle, elapsedTime float32) {
	/**/
	ball.x += ball.xv * elapsedTime
	ball.y += ball.yv * elapsedTime

	/* Screen boundaries - Top and Bottom */
	if int(ball.y-ball.radius) < 0 {
		/* Reverse velocity */
		ball.yv = -ball.yv
		/* Resolve collision */
		ball.y = ball.radius
	} else if int(ball.y+ball.radius) > winHeight {
		/* Reverse velocity */
		ball.yv = -ball.yv
		/* Resolve collision */
		ball.y = float32(winHeight) - ball.radius
	}
	/* Screen boundaries - Left */
	if int(ball.x-ball.radius) < 0 {
		/* Right paddle scores*/
		rightPaddle.score++
		/* Center ball */
		ball.pos = getCenter()
		/* Update state*/
		state = start
	}
	/* Screen boundaries - Right */
	if int(ball.x+ball.radius) > winWidth {
		/* Left paddle scores*/
		leftPaddle.score++
		/* Center ball */
		ball.pos = getCenter()
		/* Update state*/
		state = start

	}

	/* Hit a paddle left */
	if ball.x-ball.radius < leftPaddle.x+leftPaddle.w/2 {
		/**/
		if ball.y > leftPaddle.y-leftPaddle.h/2 &&
			ball.y < leftPaddle.y+leftPaddle.h/2 {
			/* Reverse velocity */
			ball.xv = -ball.xv
			/* Resolve collision */
			ball.x = leftPaddle.x + leftPaddle.w/2 + ball.radius
		}
	}
	/* Hit a paddle right */
	if ball.x+ball.radius > rightPaddle.x-rightPaddle.w/2 {
		/**/
		if ball.y > rightPaddle.y-rightPaddle.h/2 &&
			ball.y < rightPaddle.y+rightPaddle.h/2 {
			/* Reverse velocity */
			ball.xv = -ball.xv
			/* Resolve collision */
			ball.x = rightPaddle.x - rightPaddle.w/2 - ball.radius
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

/* get center of screen */
func getCenter() pos {
	/**/
	return pos{float32(winWidth / 2), float32(winHeight / 2)}
}

/* Draw number */
func drawNumber(pos pos, colour colour, size int, num int, pixels []byte) {
	/**/
	startX := int(pos.x) - (size*3)/2
	startY := int(pos.y) - (size*5)/2
	/* Draw a square*/
	for i, v := range nums[num] {
		/**/
		if v == 1 {
			/* Iterate Y*/
			for y := startY; y < startY+size; y++ {
				/* Iterate X*/
				for x := startX; x < startX+size; x++ {
					/**/
					setPixel(x, y, colour, pixels)
				}
			}
		}
		/**/
		startX += size
		/* Go to enxt square and after 3 squares go to next line*/
		if (i+1)%3 == 0 {
			/* Next line */
			startY += size
			startX -= size * 3
		}
	}

}

/* Linear interpolation*/
func lerp(a float32, b float32, pct float32) float32 {
	/* Interpolate */
	return a + pct*(b-a)
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
	window, err := sdl.CreateWindow("Popong popong", sdl.WINDOWPOS_UNDEFINED,
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

	/* Time */
	var frameStart time.Time
	var elapsedTime float32

	/*Pixels*/
	pixels := make([]byte, winWidth*winHeight*4) // 4 bytes: Alpha, Blue, Green, Red

	/*** Level 1***/
	/* Player paddles*/
	player1 := paddle{pos{50, 300}, 20, 100, 300, 0, colour{255, 255, 255}}
	player2 := paddle{pos{float32(winWidth - 50), 300}, 20, 100, 300, 0, colour{255, 255, 255}}

	/* Ball */
	ball := ball{pos{400, 300}, 20, 400, 400, colour{255, 255, 255}}

	/* Keyboard state */
	keyState := sdl.GetKeyboardState()

	/**
	 *
	 *    Game loop
	 *
	 **/
	for {
		/* Time initialise */
		frameStart = time.Now()

		/**/
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

		/* Switch state*/
		switch state {

		/* play */
		case play:
			/* Updates */
			player1.update(keyState, elapsedTime)
			player2.aiUpdate(&ball, elapsedTime)
			ball.update(&player1, &player2, elapsedTime)

		/* start */
		case start:
			/* SPACE is pressed */
			if keyState[sdl.SCANCODE_SPACE] != 0 {
				/* 10 scores */
				if player1.score == MAX_SCORE || player2.score == MAX_SCORE {
					/* Change level */
					if player1.score == MAX_SCORE {
						/* Switch level*/
						switch level {
						/**/
						case level1:
							/**/
							level = level2
							/* AI paddle */
							//player2 = paddle{pos{float32(winWidth - 50), 300}, 20, 100, 300, 0, colour{255, 255, 255}}
							player2.score = 1
							player2.speed = 500

							/* Ball */
							ball.xv = 300
							ball.yv = 300
							ball.radius = 10
							ball.colour = colour{255, 255, 255}
							/**/
						case level2:
							/**/
							level = level3
							/* AI paddle */
							//player2 = paddle{pos{float32(winWidth - 50), 300}, 20, 100, 300, 0, colour{255, 255, 255}}
							player2.score = 0
							player2.speed = 400

							/* Ball */
							ball.xv = 500
							ball.yv = 500
							ball.radius = 10
							ball.colour = colour{255, 255, 255}
							/**/
						case level3:
							/**/
							level = level4
							/* AI paddle */
							//player2 = paddle{pos{float32(winWidth - 50), 300}, 20, 100, 300, 0, colour{255, 255, 255}}
							player2.score = 0
							player2.speed = 400

							/* Ball */
							ball.xv = 500
							ball.yv = 500
							ball.radius = 5
							ball.colour = colour{255, 255, 255}
							/**/
						case level4:
							level = level5
							/**/
						case level5:
							level = level1
						}
					} else if player2.score == MAX_SCORE {
						/* GAME OVER*/
						state = gameover
					}
					/* Set scores */
					player1.score = 0
					player2.score = 0
				}
				/* Change game state */
				if level != level5 && state != gameover {
					state = play
				} else {
					state = finished
				}
			}

			/* gameover */
		case gameover:
			/* Clear screen */
			clear(pixels)

			/**/
			for true {
				if keyState[sdl.SCANCODE_SPACE] != 0 {
					state = start
					break
				}
			}

			/* finished */

		}

		/* Clear screen */
		clear(pixels)

		/* Draw paddles*/
		player1.draw(pixels)
		player2.draw(pixels)
		/* Draw ball */
		ball.draw(pixels)
		/* Draw level */
		drawNumber(pos{float32(winWidth / 2), 35}, colour{255, 255, 255}, 10, (int(level) + 1), pixels)

		/*Update texture*/
		texture.Update(nil, pixels, winWidth*4)
		renderer.Copy(texture, nil, nil)
		renderer.Present()

		/* Time - To keep framerate for every system*/
		elapsedTime = float32(time.Since(frameStart).Seconds())
		/* Limit framerate */
		if elapsedTime < MIN_ELAPSED_TIME {
			/**/
			sdl.Delay(5 - uint32(elapsedTime*1000.0))
			elapsedTime = float32(time.Since(frameStart).Seconds())
		}
	}
}
