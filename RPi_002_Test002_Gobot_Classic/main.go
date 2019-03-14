/*
  Classic Gobot

 The simplest way to create robots, and drones, and Internet
connected things, is to use "Classic Gobot".
*/
/*
For extended PWM support on the Raspberry Pi, you will need to use a program called pi-blaster. You can follow the instructions for pi-blaster install in the pi-blaster repo here:

https://github.com/sarfata/pi-blaster
*/
/*
 Build using W10:

$env:GOOS="linux"
$env:GOARCH="arm"
$env:GOARM="6"
go build -o <name>

arm 6 is for Pi Zero
replace <name> for the name of the output file
*/
package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()
	led := gpio.NewLedDriver(r, "7")

	work := func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{r},
		[]gobot.Device{led},
		work,
	)

	robot.Start()
}
