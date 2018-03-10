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

	fmt.Println("Press a button a-h (vertical) or 1-8 (horizontal) on the launchpad")

	pad.Clear()
	f := padx.Widgets8x8()
	ch := pad.Listen()
	for {
		hit := <-ch
		pad.Clear()

		// check if the key is a menu button : a-h, 1-8
		if btn := padx.Get(hit); btn != nil {
			// it's a 'button'

			// display the key pressed
			fmt.Println("pressed:", hit)

			f.Widgets[btn.Name].Blink(pad, padx.ColorRedFull, padx.ColorOff, 100*time.Millisecond, 15)
		} else {
			fmt.Println("pressed (not a button):", hit)
		}
	}
}
