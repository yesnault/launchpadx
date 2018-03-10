package main

import (
	"fmt"
	"time"

	"github.com/rakyll/launchpad"

	"github.com/yesnault/launchpadx/widgets"
)

func main() {
	pad, err := launchpad.Open()
	if err != nil {
		fmt.Printf("error while openning connection to launchpad: %v", err)
	}
	defer pad.Close()
	pad.Clear()

	// draw 4 squares, "Russian dolls"
	widgets.Rectangle{P: launchpad.Hit{X: 0, Y: 0}, Width: 8, Height: 8}.Paint(pad, widgets.ColorGreenFull, 10*time.Millisecond)
	widgets.Rectangle{P: launchpad.Hit{X: 1, Y: 1}, Width: 6, Height: 6}.Paint(pad, widgets.ColorRedFull, 20*time.Millisecond)
	widgets.Rectangle{P: launchpad.Hit{X: 2, Y: 2}, Width: 4, Height: 4}.Paint(pad, widgets.ColorYellowFull, 50*time.Millisecond)
	widgets.Rectangle{P: launchpad.Hit{X: 3, Y: 3}, Width: 2, Height: 2}.Paint(pad, widgets.ColorAmberFull, 100*time.Millisecond)

	time.Sleep(1 * time.Second)

	fmt.Println("Then clear all with 15ms transition between each pixel")
	widgets.Rectangle{P: launchpad.Hit{X: 0, Y: 0}, Width: 8, Height: 8}.Clear(pad, 15*time.Millisecond)

	time.Sleep(1 * time.Second)

	// and move a rectangle
	r := widgets.Rectangle{P: launchpad.Hit{X: 0, Y: 0}, Width: 2, Height: 2}
	r.Paint(pad, widgets.ColorGreenFull, 10*time.Millisecond)

	time.Sleep(300 * time.Millisecond)

	r.ScrollTo(pad, launchpad.Hit{X: 6, Y: 6}, widgets.ColorGreenFull, 300*time.Millisecond)
	time.Sleep(300 * time.Millisecond)

	r.ScrollTo(pad, launchpad.Hit{X: 1, Y: 3}, widgets.ColorYellowFull, 300*time.Millisecond)
	time.Sleep(300 * time.Millisecond)

	r.ScrollTo(pad, launchpad.Hit{X: 5, Y: 3}, widgets.ColorYellowFull, 300*time.Millisecond)
	time.Sleep(300 * time.Millisecond)

	r.ScrollTo(pad, launchpad.Hit{X: 3, Y: 3}, widgets.ColorRedFull, 300*time.Millisecond)

	time.Sleep(300 * time.Millisecond)
}
