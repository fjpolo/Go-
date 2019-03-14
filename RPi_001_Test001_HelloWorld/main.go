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
	"fmt"  // print
	"time" // OS time

	"github.com/stianeikeland/go-rpio" // Raspberry Pi GPIOs
)

/*

	Does not support Raspberry Pi Zero!!!!

*/

func main() {
	fmt.Println("Opening gpio")
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("Unable to open gpio", err.Error()))
	}

	/* Defer to close when main is done */
	defer rpio.Close()

	/* Create pin instance */
	pin := rpio.Pin(18)
	/* Configure pin 18 as output */
	pin.Output()

	/* loop 20 times */
	for x := 0; x < 20; x++ {
		/* Toggle */
		pin.Toggle()
		/* Sleep 500mS */
		time.Sleep(time.Second / 2)
	}
}
