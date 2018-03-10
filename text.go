package padx

import (
	"github.com/rakyll/launchpad"
)

// Text returns a new text
func Text(text string, direction Direction, font Font, offsetX, offsetY int) Custom {
	c := Custom{
		OffsetX: offsetX,
		OffsetY: offsetY,
		Hits:    []launchpad.Hit{},
	}

	var width int
	var height int
	for _, r := range text {
		for _, h := range font.Widgets[r].Hits {
			h.X = width + h.X
			c.Hits = append(c.Hits, h)
		}
		width += font.Widgets[r].Width
		height = font.Widgets[r].Height
	}

	c.Width = width
	c.Height = height

	return c
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
