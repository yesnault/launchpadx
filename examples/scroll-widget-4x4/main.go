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

	w := padx.Widgets4x4().Widgets['h']

	w.Paint(pad, padx.ColorGreenFull, 0)
	time.Sleep(300 * time.Millisecond)

	w.ScrollTo(pad, launchpad.Hit{X: 5, Y: 4}, padx.ColorGreenFull, 300*time.Millisecond)

	time.Sleep(1 * time.Second)
}
