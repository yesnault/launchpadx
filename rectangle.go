package padx

import (
	"github.com/rakyll/launchpad"
)

// Rectangle returns a new custom widget
func Rectangle(posX, posY, width, height int) Custom {
	c := Custom{
		OffsetX: 0,
		OffsetY: 0,
		Width:   width,
		Height:  height,
		Hits:    []launchpad.Hit{},
	}

	for x := 0; x < width; x++ {
		for y := 0; y < width; y++ {
			c.Hits = append(c.Hits, launchpad.Hit{X: posX + x, Y: posY + y})
		}
	}

	return c
}
