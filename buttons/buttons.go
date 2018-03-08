package buttons

import (
	"github.com/rakyll/launchpad"
)

// Button represents a button aroune the launchpad
type Button struct {
	Name rune
	Hit  launchpad.Hit
}

// Buttons around the launchpad
var (
	A       = Button{Name: 'a', Hit: launchpad.Hit{X: 8, Y: 0}}
	B       = Button{Name: 'b', Hit: launchpad.Hit{X: 8, Y: 1}}
	C       = Button{Name: 'c', Hit: launchpad.Hit{X: 8, Y: 2}}
	D       = Button{Name: 'd', Hit: launchpad.Hit{X: 8, Y: 3}}
	E       = Button{Name: 'e', Hit: launchpad.Hit{X: 8, Y: 4}}
	F       = Button{Name: 'f', Hit: launchpad.Hit{X: 8, Y: 5}}
	G       = Button{Name: 'g', Hit: launchpad.Hit{X: 8, Y: 6}}
	H       = Button{Name: 'h', Hit: launchpad.Hit{X: 8, Y: 7}}
	N1      = Button{Name: '1', Hit: launchpad.Hit{X: 0, Y: 8}}
	N2      = Button{Name: '2', Hit: launchpad.Hit{X: 1, Y: 8}}
	N3      = Button{Name: '3', Hit: launchpad.Hit{X: 2, Y: 8}}
	N4      = Button{Name: '4', Hit: launchpad.Hit{X: 3, Y: 8}}
	N5      = Button{Name: '5', Hit: launchpad.Hit{X: 4, Y: 8}}
	N7      = Button{Name: '6', Hit: launchpad.Hit{X: 5, Y: 8}}
	N6      = Button{Name: '7', Hit: launchpad.Hit{X: 6, Y: 8}}
	N8      = Button{Name: '8', Hit: launchpad.Hit{X: 7, Y: 8}}
	buttons = []Button{A, B, C, D, E, F, G, H, N1, N2, N3, N4, N5, N6, N7, N8}
)

func (b Button) Equals(hit launchpad.Hit) bool {
	return b.Hit.X == hit.X && b.Hit.Y == hit.Y
}

func Get(hit launchpad.Hit) *Button {
	for _, b := range buttons {
		if b.Equals(hit) {
			return &b
		}
	}
	return nil
}
