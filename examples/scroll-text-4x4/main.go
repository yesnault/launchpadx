package main

import (
	"fmt"
	"time"

	"github.com/rakyll/launchpad"

	padx "github.com/yesnault/launchpadx"
)

func main() {
	pad, err := launchpad.Open()
	if err != nil {
		fmt.Printf("error while openning connection to launchpad: %v", err)
	}
	defer pad.Close()

	fmt.Println("Press a button a-h (vertical) or 1-8 (horizontal) on the launchpad")

	pad.Clear()
	ch := pad.Listen()
	text := padx.NewText("foobar", padx.DirectionLeftToRight, padx.Widgets4x4())
	for {
		hit := <-ch
		pad.Clear()

		// check if the key is a menu button : a-h, 1-8
		if btn := padx.Get(hit); btn != nil {
			// it's a 'button'

			// display the key pressed
			fmt.Println("pressed:", hit)

			// scroll horizontally, right to left
			text.Scroll(pad, padx.ColorRedFull, padx.DirectionRightToLeft, 100*time.Millisecond)
		} else {
			fmt.Println("pressed (not a button):", hit)
		}
	}
}
