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
		fmt.Printf("error while opening connection to launchpad: %v", err)
	}
	defer pad.Close()
	pad.Clear()

	// draw 4 squares, "Russian dolls"
	padx.Rectangle(0, 0, 8, 8).Paint(pad, padx.ColorGreenFull, 10*time.Millisecond)
	padx.Rectangle(1, 1, 6, 6).Paint(pad, padx.ColorRedFull, 20*time.Millisecond)
	padx.Rectangle(2, 2, 4, 4).Paint(pad, padx.ColorYellowFull, 50*time.Millisecond)
	padx.Rectangle(3, 3, 2, 2).Paint(pad, padx.ColorAmberFull, 100*time.Millisecond)

	time.Sleep(1 * time.Second)

	fmt.Println("Then clear all with 15ms transition between each pixel")
	padx.Rectangle(0, 0, 8, 8).Clear(pad, 15*time.Millisecond)

	time.Sleep(1 * time.Second)

	// and move a rectangle
	r := padx.Rectangle(0, 0, 2, 2)
	r.Paint(pad, padx.ColorGreenFull, 10*time.Millisecond)

	time.Sleep(300 * time.Millisecond)

	r.ScrollTo(pad, 6, 6, padx.ColorGreenFull, 300*time.Millisecond)
	time.Sleep(300 * time.Millisecond)

	r.ScrollTo(pad, 1, 3, padx.ColorYellowFull, 300*time.Millisecond)
	time.Sleep(300 * time.Millisecond)

	r.ScrollTo(pad, 5, 3, padx.ColorYellowFull, 300*time.Millisecond)
	time.Sleep(300 * time.Millisecond)

	r.ScrollTo(pad, 3, 3, padx.ColorRedFull, 300*time.Millisecond)

	time.Sleep(300 * time.Millisecond)
}
