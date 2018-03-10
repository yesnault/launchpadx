package padx

import (
	"time"

	"github.com/rakyll/launchpad"
)

// Font represents a set of widgets
type Font struct {
	Widgets map[rune]Custom
}

// Widget have to be implemented by widget
type Widget interface {
	Paint(pad *launchpad.Launchpad, color Color, d time.Duration)
	Clear(pad *launchpad.Launchpad, d time.Duration)
	Move(pad *launchpad.Launchpad, to launchpad.Hit, color Color, d time.Duration)
}
