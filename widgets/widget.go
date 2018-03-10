package widgets

import (
	"time"

	"github.com/rakyll/launchpad"
)

// Widget have to be implemented by widget
type Widget interface {
	Paint(pad *launchpad.Launchpad, color Color, d time.Duration)
	Clear(pad *launchpad.Launchpad, d time.Duration)
	ScrollTo(pad *launchpad.Launchpad, to launchpad.Hit, color Color, d time.Duration)
}

// Custom represents a custom widget and all its hits (x, y) for behing paint
type Custom struct {
	OffsetX int
	OffsetY int
	Width   int
	Height  int
	Hits    []launchpad.Hit
}

// Size of the font definition
type Size struct {
	Height int
	Width  int
}

// Color represent a color on launchpad mini: green and red
type Color struct {
	Name  string
	Green int
	Red   int
}

// Colors
var (
	ColorOff         = Color{Name: "off", Green: 0, Red: 0}
	ColorGreenLow    = Color{Name: "greenLow", Green: 1, Red: 0}
	ColorGreenMedium = Color{Name: "greenMedium", Green: 2, Red: 0}
	ColorGreenFull   = Color{Name: "greenFull", Green: 3, Red: 0}
	ColorYellowFull  = Color{Name: "yellowFull", Green: 3, Red: 1}
	ColorRedLight    = Color{Name: "redLight", Green: 1, Red: 2}
	ColorOrangeLight = Color{Name: "orangeLight", Green: 3, Red: 2}
	ColorRedFull     = Color{Name: "redFull", Green: 0, Red: 3}
	ColorAmberFull   = Color{Name: "amberFull", Green: 2, Red: 3}
	ColorOrangeFull  = Color{Name: "orangeFull", Green: 3, Red: 3}

	Colors = []Color{
		ColorGreenLow,
		ColorGreenMedium,
		ColorGreenFull,
		ColorYellowFull,
		ColorRedLight,
		ColorOrangeLight,
		ColorRedFull,
		ColorAmberFull,
		ColorOrangeFull,
	}
)
