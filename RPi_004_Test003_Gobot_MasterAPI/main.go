/*
  Master Gobot

Use Master Gobot when you want to provide an API to your robot,
drone, or Internet connected thing.

Also, use Master Gobot when you want to create a swarm of devices,
and control them as a single unit.
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
	"gobot.io/x/gobot/api"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	master := gobot.NewMaster()
	a := api.NewAPI(master)
	a.Start()

	raspiAdaptor := raspi.NewAdaptor()
	led := gpio.NewLedDriver(raspiAdaptor, "13")

	work := func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("R2D2",
		[]gobot.Connection{raspiAdaptor},
		[]gobot.Device{led},
		work,
	)

	master.AddRobot(robot)

	master.Start()
}
