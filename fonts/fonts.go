package fonts

import (
	"time"

	"github.com/rakyll/launchpad"
)

// Size of the font definition
type Size struct {
	Height int
	Width  int
}

// Font represents a font, with a size and charset
type Font struct {
	Size    Size
	CharSet map[string]Character
}

// Character represents a characters and all its hits (x, y) for behing paint
type Character struct {
	OffsetX int
	OffsetY int
	Hits    []launchpad.Hit
}

// Color represent a color on launchpad mini: green and red
type Color struct {
	Name  string
	Green int
	Red   int
}

var (
	Off         = Color{Name: "off", Green: 0, Red: 0}
	GreenLow    = Color{Name: "greenLow", Green: 1, Red: 0}
	GreenMedium = Color{Name: "greenMedium", Green: 2, Red: 0}
	GreenFull   = Color{Name: "greenFull", Green: 3, Red: 0}
	YellowFull  = Color{Name: "yellowFull", Green: 3, Red: 1}
	RedLight    = Color{Name: "redLight", Green: 1, Red: 2}
	OrangeLight = Color{Name: "orangeLight", Green: 3, Red: 2}
	RedFull     = Color{Name: "redFull", Green: 0, Red: 3}
	AmberFull   = Color{Name: "amberFull", Green: 2, Red: 3}
	OrangeFull  = Color{Name: "orangeFull", Green: 3, Red: 3}

	Colors = []Color{GreenLow, GreenMedium, GreenFull, YellowFull, RedLight, OrangeLight, RedFull, AmberFull, OrangeFull}
)

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

// Paint colors a character on launchpad
func (c Character) Paint(pad *launchpad.Launchpad, color Color) {
	for _, h := range c.Hits {
		pad.Light(h.X+c.OffsetX, h.Y+c.OffsetY, color.Green, color.Red)
	}
}

// Blink blinks a character with two colors
// duration int     Duration of transition between colorA and colorB
// repeats int      Number of repetitions
func (c Character) Blink(pad *launchpad.Launchpad, colorA, colorB Color, duration time.Duration, repeats int) {
	for r := 0; r < repeats; r++ {
		if r%2 == 0 {
			c.Paint(pad, colorA)
		} else {
			c.Paint(pad, colorB)
		}
		time.Sleep(duration)
	}
}
