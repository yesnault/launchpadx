package main

import (
	"fmt"
	"math/rand"

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
	f := padx.Widgets8x8()
	ch := pad.Listen()
	for {
		hit := <-ch
		pad.Clear()

		// check if the key is a menu button : a-h, 1-8
		if btn := padx.Get(hit); btn != nil {
			// it's a 'button'

			// get a random color
			color := padx.Colors[rand.Intn(len(padx.Colors))]

			// display the key pressed
			fmt.Println("pressed:", hit, "with color:", color.Name)

			f.Widgets[btn.Name].Paint(pad, color, 0)
		} else {
			fmt.Println("pressed (not a button):", hit)
		}
	}
}
