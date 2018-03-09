package main

import (
	"fmt"
	"math/rand"

	"github.com/rakyll/launchpad"

	"github.com/yesnault/launchpadx/buttons"
	"github.com/yesnault/launchpadx/fonts"
	"github.com/yesnault/launchpadx/widgets"
)

func main() {
	pad, err := launchpad.Open()
	if err != nil {
		fmt.Printf("error while openning connection to launchpad: %v", err)
	}
	defer pad.Close()

	fmt.Println("Press a button a-h (vertical) or 1-8 (horizontal) on the launchpad")

	pad.Clear()
	f := fonts.Size4x4()
	ch := pad.Listen()
	for {
		hit := <-ch
		pad.Clear()

		// check if the key is a menu button : a-h, 1-8
		if btn := buttons.Get(hit); btn != nil {
			// it's a 'button'

			// get a random color
			color := widgets.Colors[rand.Intn(len(widgets.Colors))]

			// display the key pressed
			fmt.Println("pressed:", hit, "with color:", color.Name)

			f.CharSet[btn.Name].Paint(pad, color)
		} else {
			fmt.Println("pressed (not a button):", hit)
		}
	}
}
