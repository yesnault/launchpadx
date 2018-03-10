package padx

import (
	"time"

	"github.com/rakyll/launchpad"
)

// Font represents a set of widgets
type Font struct {
	Widgets map[rune]Custom
}

// Text represents a list of Widgets
type Text struct {
	OffsetX   int
	OffsetY   int
	Text      string
	Font      Font
	Direction Direction
	Widgets   []Custom
}

// NewText returns a new text
func NewText(text string, direction Direction, font Font) Text {
	t := Text{
		Text:      text,
		Direction: direction,
		Font:      font,
		Widgets:   make([]Custom, len(text)),
	}
	for i, r := range text {
		t.Widgets[i] = font.Widgets[r]
	}
	return t
}

// Paint colors a text on launchpad
func (t *Text) Paint(pad *launchpad.Launchpad, color Color) {
	for pos, w := range t.Widgets {
		switch t.Direction {
		case DirectionLeftToRight:
			w.OffsetX = t.OffsetX + pos*w.Width
		case DirectionRightToLeft:
			w.OffsetX = t.OffsetX - pos*w.Width
		case DirectionTopToBottom:
			w.OffsetY = t.OffsetY - pos*w.Height
		case DirectionBottomToTop:
			w.OffsetY = t.OffsetY + pos*w.Height
		}
		w.Paint(pad, color, 0)
	}
}

// Scroll scrolls a text
// duration int     Duration of transition
func (t Text) Scroll(pad *launchpad.Launchpad, color Color, direction Direction, duration time.Duration) {
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

	for r := 0; r < len(t.Widgets)*t.Widgets[0].Width*2+16; r++ {
		if r%2 == 0 {
			t.OffsetX += xoff
			t.OffsetY += yoff
			t.Paint(pad, color)
			time.Sleep(duration)
		} else {
			t.Paint(pad, ColorOff)
		}

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
