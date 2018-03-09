package fonts

import (
	"time"

	"github.com/yesnault/launchpadx/widgets"

	"github.com/rakyll/launchpad"
)

// Font represents a font, with a size and charset
type Font struct {
	Size    widgets.Size
	CharSet map[rune]Character
}

// Character represents a characters and all its hits (x, y) for behing paint
type Character struct {
	OffsetX int
	OffsetY int
	Hits    []launchpad.Hit
}

// Text represents a list of characters
type Text struct {
	OffsetX    int
	OffsetY    int
	Text       string
	Font       Font
	Direction  Direction
	Characters []Character
}

func process(lines []string) []launchpad.Hit {
	hits := []launchpad.Hit{}
	for y, line := range lines {
		for x, pixel := range line {
			if pixel == '0' {
				hits = append(hits, launchpad.Hit{X: x, Y: y})
			}
		}
	}
	return hits
}

// NewText returns a new text
func NewText(text string, direction Direction, font Font) Text {
	t := Text{
		Text:       text,
		Direction:  direction,
		Font:       font,
		Characters: make([]Character, len(text)),
	}
	for i, r := range text {
		t.Characters[i] = font.CharSet[r]
	}
	return t
}

// Paint colors a text on launchpad
func (t *Text) Paint(pad *launchpad.Launchpad, color widgets.Color) {
	for pos, c := range t.Characters {
		switch t.Direction {
		case DirectionLeftToRight:
			c.OffsetX = t.OffsetX + pos*t.Font.Size.Width
		case DirectionRightToLeft:
			c.OffsetX = t.OffsetX - pos*t.Font.Size.Width
		case DirectionTopToBottom:
			c.OffsetY = t.OffsetY - pos*t.Font.Size.Height
		case DirectionBottomToTop:
			c.OffsetY = t.OffsetY + pos*t.Font.Size.Height
		}
		c.Paint(pad, color)
	}
}

// Scroll scrolls a text
// duration int     Duration of transition
func (t Text) Scroll(pad *launchpad.Launchpad, color widgets.Color, direction Direction, duration time.Duration) {
	var xoff, yoff int
	switch direction {
	case DirectionRightToLeft:
		t.OffsetX = 8
		xoff = -1
	case DirectionLeftToRight:
		t.OffsetX = -8
		xoff = 1
	case DirectionBottomToTop:
		t.OffsetY = 8
		yoff = -1
	case DirectionTopToBottom:
		t.OffsetY = -8
		yoff = 1
	}

	for r := 0; r < len(t.Characters)*t.Font.Size.Width*2+16; r++ {
		if r%2 == 0 {
			t.OffsetX += xoff
			t.OffsetY += yoff
			t.Paint(pad, color)
			time.Sleep(duration)
		} else {
			t.Paint(pad, widgets.ColorOff)
		}

	}
}

// Paint colors a character on launchpad
func (c Character) Paint(pad *launchpad.Launchpad, color widgets.Color) {
	for _, h := range c.Hits {
		if h.X+c.OffsetX >= 0 && h.X+c.OffsetX < 8 && h.Y+c.OffsetY >= 0 && h.Y+c.OffsetY < 8 {
			pad.Light(h.X+c.OffsetX, h.Y+c.OffsetY, color.Green, color.Red)
		}
	}
}

// Blink blinks a character with two colors
// duration int     Duration of transition between colorA and colorB
// repeats int      Number of repetitions
func (c Character) Blink(pad *launchpad.Launchpad, colorA, colorB widgets.Color, duration time.Duration, repeats int) {
	for r := 0; r < repeats; r++ {
		if r%2 == 0 {
			c.Paint(pad, colorA)
		} else {
			c.Paint(pad, colorB)
		}
		time.Sleep(duration)
	}
}

// Direction represents a direction for scrolling
type Direction int

// direction for scroll
const (
	DirectionRightToLeft Direction = iota
	DirectionLeftToRight
	DirectionTopToBottom
	DirectionBottomToTop
)

// Scroll scrolls a character
// duration int     Duration of transition
func (c Character) Scroll(pad *launchpad.Launchpad, color widgets.Color, direction Direction, duration time.Duration) {
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
			c.Paint(pad, color)
			time.Sleep(duration)
		} else {
			c.Paint(pad, widgets.ColorOff)
		}
	}
}
