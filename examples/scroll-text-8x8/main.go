package main

import (
	"fmt"
	"time"

	"github.com/rakyll/launchpad"
	"github.com/yesnault/launchpadx/buttons"
	"github.com/yesnault/launchpadx/fonts"
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
	text := fonts.NewText("foobar", fonts.DirectionLeftToRight, fonts.Size8x8())
	for {
		hit := <-ch
		pad.Clear()

		// check if the key is a menu button : a-h, 1-8
		if btn := buttons.Get(hit); btn != nil {
			// it's a 'button'

			// display the key pressed
			fmt.Println("pressed:", hit)

			// scroll horizontally, right to left
			text.Scroll(pad, fonts.RedFull, fonts.DirectionRightToLeft, 100*time.Millisecond)
		} else {
			fmt.Println("pressed (not a button):", hit)
		}
	}
}
