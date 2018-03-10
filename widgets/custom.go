package widgets

import (
	"time"

	"github.com/rakyll/launchpad"
)

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
func (c Custom) Paint(pad *launchpad.Launchpad, color Color, d time.Duration) {
	for _, h := range c.Hits {
		if h.X+c.OffsetX >= 0 && h.X+c.OffsetX < 8 && h.Y+c.OffsetY >= 0 && h.Y+c.OffsetY < 8 {
			pad.Light(h.X+c.OffsetX, h.Y+c.OffsetY, color.Green, color.Red)
		}
		time.Sleep(d)
	}
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
func (c Custom) ScrollTo(pad *launchpad.Launchpad, to launchpad.Hit, color Color, d time.Duration) {

}

// Scroll scrolls a custom widget
// duration int     Duration of transition
func (c Custom) Scroll(pad *launchpad.Launchpad, color Color, direction Direction, duration time.Duration) {
	var xoff, yoff int
	switch direction {
	case DirectionRightToLeft:
		c.OffsetX = 8
		xoff = -1
	case DirectionLeftToRight:
		c.OffsetX = -8
		xoff = 1
	case DirectionBottomToTop:
		c.OffsetY = 8
		yoff = -1
	case DirectionTopToBottom:
		c.OffsetY = -8
		yoff = 1
	}

	for r := 0; r < 28; r++ {
		if r%2 == 0 {
			c.OffsetX += xoff
			c.OffsetY += yoff
			c.Paint(pad, color, 0)
			time.Sleep(duration)
		} else {
			c.Paint(pad, ColorOff, 0)
		}
	}
}
