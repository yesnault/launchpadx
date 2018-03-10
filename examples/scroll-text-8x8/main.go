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

	text := padx.Text("foobar", padx.DirectionLeftToRight, padx.Widgets8x8(), 8, 0)
	pad.Clear()

	text.Paint(pad, padx.ColorRedFull, 0)
	time.Sleep(100 * time.Millisecond)

	// from right to left
	text.ScrollTo(pad, -text.Width, 0, padx.ColorRedFull, 100*time.Millisecond)

	// from left to right
	text.ScrollTo(pad, text.Width, 0, padx.ColorRedFull, 100*time.Millisecond)

	time.Sleep(100 * time.Millisecond)
}
