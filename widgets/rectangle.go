package widgets

import (
	"time"

	"github.com/rakyll/launchpad"
)

// Rectangle represents a Rectangle on the launchpad
type Rectangle struct {
	P      launchpad.Hit
	Width  int
	Height int
}

// Paint paint a rectangle
func (r Rectangle) Paint(pad *launchpad.Launchpad, color Color, d time.Duration) {
	for x := r.P.X; x < r.P.X+r.Width; x++ {
		for y := r.P.Y; y < r.P.Y+r.Height; y++ {
			pad.Light(x, y, color.Green, color.Red)
			time.Sleep(d)
		}
	}
}
