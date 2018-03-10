package padx

import (
	"time"

	"github.com/rakyll/launchpad"
)

// Custom represents a custom widget and all its hits (x, y) for behing paint
type Custom struct {
	OffsetX int
	OffsetY int
	Width   int
	Height  int
	Hits    []launchpad.Hit
}

// NewCustom initializes a custom widget
func NewCustom(lines []string) Custom {
	hits := []launchpad.Hit{}
	for y, line := range lines {
		for x, pixel := range line {
			if pixel == '0' {
				hits = append(hits, launchpad.Hit{X: x, Y: y})
			}
		}
	}

	return Custom{
		Width:  len(lines[0]),
		Height: len(lines),
		Hits:   hits,
	}
}

// Paint colors a custom widget on launchpad
func (c Custom) Paint(pad *launchpad.Launchpad, color Color, d time.Duration) []launchpad.Hit {
	hits := []launchpad.Hit{}
	for _, h := range c.Hits {
		x := h.X + c.OffsetX
		y := h.Y + c.OffsetY
		if x >= 0 && x < 8 && y >= 0 && y < 8 {
			pad.Light(x, y, color.Green, color.Red)
			hits = append(hits, launchpad.Hit{X: x, Y: y})
			time.Sleep(d)
		}
	}
	return hits
}

// Clear paint a custom widget with colorOff
func (c Custom) Clear(pad *launchpad.Launchpad, d time.Duration) {
	c.Paint(pad, ColorOff, d)
}

// Blink blinks a custom widget with two colors
// duration int     Duration of transition between colorA and colorB
// repeats int      Number of repetitions
func (c Custom) Blink(pad *launchpad.Launchpad, colorA, colorB Color, duration time.Duration, repeats int) {
	for r := 0; r < repeats; r++ {
		if r%2 == 0 {
			c.Paint(pad, colorA, 0)
		} else {
			c.Paint(pad, colorB, 0)
		}
		time.Sleep(duration)
	}
}

// ScrollTo scroll custom widget to another position
func (c *Custom) ScrollTo(pad *launchpad.Launchpad, toX, toY int, color Color, d time.Duration) {
	for m := 1; m <= c.Width*3; m++ {
		newPos := launchpad.Hit{X: c.OffsetX, Y: c.OffsetY}
		if toX > c.OffsetX {
			newPos.X = c.OffsetX + 1
		} else if toX < c.OffsetX {
			newPos.X = c.OffsetX - 1
		}

		if toY > c.OffsetY {
			newPos.Y = c.OffsetY + 1
		} else if toY < c.OffsetY {
			newPos.Y = c.OffsetY - 1
		}

		originalX := c.OffsetX
		originalY := c.OffsetY

		c.OffsetX = newPos.X
		c.OffsetY = newPos.Y
		hits := c.Paint(pad, color, 0)

		for x := originalX; x < originalX+c.Width; x++ {
			for y := originalY; y < originalY+c.Height; y++ {
				var isOn bool
				for _, hit := range hits {
					if hit.X == x && hit.Y == y {
						// do not off this led
						isOn = true
					}
				}
				if !isOn && x >= 0 && x < 8 && y >= 0 && y < 8 {
					pad.Light(x, y, 0, 0)
				}

			}
		}
		time.Sleep(d)
		c.OffsetX = newPos.X
		c.OffsetY = newPos.Y
		if c.OffsetX == toX && c.OffsetY == toY {
			// destination reached
			break
		}
	}
	c.OffsetX = toX
	c.OffsetY = toY
}
