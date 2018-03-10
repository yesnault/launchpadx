package widgets

import (
	"time"

	"github.com/rakyll/launchpad"
)

// Rectangle represents a Rectangle on the launchpad
type Rectangle struct {
	P      launchpad.Hit
	Color  Color
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

// Clear paint a rectangle with colorOff
func (r Rectangle) Clear(pad *launchpad.Launchpad, d time.Duration) {
	r.Paint(pad, ColorOff, d)
}

// ScrollTo scroll rectangle to another position
func (r *Rectangle) ScrollTo(pad *launchpad.Launchpad, to launchpad.Hit, color Color, d time.Duration) {
	for m := 1; m <= 8; m++ {
		newPos := launchpad.Hit{X: r.P.X, Y: r.P.Y}
		if to.X > r.P.X {
			newPos.X = r.P.X + 1
		} else if to.X < r.P.X {
			newPos.X = r.P.X - 1
		}

		if to.Y > r.P.Y {
			newPos.Y = r.P.Y + 1
		} else if to.Y < r.P.Y {
			newPos.Y = r.P.Y - 1
		}

		hits := []launchpad.Hit{}
		for x := newPos.X; x < newPos.X+r.Width; x++ {
			for y := newPos.Y; y < newPos.Y+r.Height; y++ {
				pad.Light(x, y, color.Green, color.Red)
				hits = append(hits, launchpad.Hit{X: x, Y: y})
			}
		}
		for x := r.P.X; x < r.P.X+r.Width; x++ {
			for y := r.P.Y; y < r.P.Y+r.Height; y++ {
				var isOn bool
				for _, hit := range hits {
					if hit.X == x && hit.Y == y {
						// do not off this led
						isOn = true
					}
				}
				if !isOn {
					pad.Light(x, y, 0, 0)
				}

			}
		}
		time.Sleep(d)
		r.P = newPos
		if r.P.X == to.X && r.P.Y == to.Y {
			// destination reached
			break
		}
	}
	r.P = to
}
