/*
  Metal Gobot

 Use Metal Gobot when you want to use the individual Gobot packages
yourself to have the greatest control, or to more easily integrate
Gobot functionality into your existing Golang programs.
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

	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	/* New Raspberry Pi Adaptor instance */
	e := raspi.NewAdaptor()
	/* Conncet adaptor */
	e.Connect()

	/* New GPIO Driver 13 instance*/
	led := gpio.NewLedDriver(e, "13")
	/* Start the driver */
	led.Start()

	for {
		/* Toggle GPIO 13 */
		led.Toggle()
		/* Sleep 1S */
		time.Sleep(1000 * time.Millisecond)
	}
}
