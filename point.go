package padx

import (
	"github.com/rakyll/launchpad"
)

// Point returns a new custom widget
func Point(posX, posY int) Custom {
	return Custom{
		OffsetX: 0,
		OffsetY: 0,
		Width:   0,
		Height:  0,
		Hits:    []launchpad.Hit{{X: posX, Y: posY}},
	}
}
