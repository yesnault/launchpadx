package main

import (
	"fmt"
	"time"

	"github.com/rakyll/launchpad"

	"github.com/yesnault/launchpadx"
)

func main() {
	pad, err := launchpad.Open()
	if err != nil {
		fmt.Printf("error while openning connection to launchpad: %v", err)
	}
	defer pad.Close()
	pad.Clear()

	padx.Point(0, 0).Paint(pad, padx.ColorGreenFull, 10*time.Millisecond)
	padx.Point(5, 3).Paint(pad, padx.ColorRedFull, 10*time.Millisecond)
	padx.Point(7, 4).Paint(pad, padx.ColorYellowFull, 10*time.Millisecond)
	padx.Point(7, 7).Paint(pad, padx.ColorAmberFull, 10*time.Millisecond)

	time.Sleep(1 * time.Second)

}
