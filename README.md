Launchpadx is a package which helps you write text / color on your Novation Launchpad Mini.

This package uses https://github.com/rakyll/launchpad


## Usage

### Character
```go

package main

import (
	"fmt"
	"math/rand"

	"github.com/rakyll/launchpad"

	"github.com/yesnault/launchpadx/buttons"
	"github.com/yesnault/launchpadx/widgets"
)


func main() {
	pad, err := launchpad.Open()
	if err != nil {
		fmt.Printf("error while openning connection to launchpad: %v", err)
	}
	defer pad.Close()

	fmt.Println("Press a button a-h (vertical) or 1-8 (horizontal) on the launchpad")

	pad.Clear()
	f := widgets.Widgets8x8()
	ch := pad.Listen()
	for {
		hit := <-ch
		pad.Clear()

		// check if the key is a menu button : a-h, 1-8
		if btn := buttons.Get(hit); btn != nil {
			// it's a 'button'

			// get a random color
			color := widgets.Colors[rand.Intn(len(widgets.Colors))]

			// display the key pressed
			fmt.Println("pressed:", hit, "with color:", color.Name)

			f.Widgets[btn.Name].Paint(pad, color, 0)
		} else {
			fmt.Println("pressed (not a button):", hit)
		}
	}
}

```

### Draw a Text

```go
[...]

text := widgets.NewText("foobar", widgets.DirectionLeftToRight, widgets.Widgets8x8())
text.Scroll(pad, widgets.ColorRedFull, widgets.DirectionRightToLeft, 100*time.Millisecond)

[...]
```


## Draw a character

```go
// initialize the font
f := widgets.Widgets8x8()

// paint a character
f.Widgets["a"].Paint(pad, widgets.ColorGreenFull)

// blink a character, 100ms transition, 15 repeats between two colors
f.Widgets["b"].Blink(pad, widgets.ColorRedFull, widgets.ColorOff, 100*time.Millisecond, 15)

// Scroll a character, 500ms transition
f.Widgets["c"].Scroll(pad, widgets.ColorRedFull, widgets.DirectionTopToBottom, 500*time.Millisecond)

```

## Draw a rectangle

```go
r := widgets.Rectangle{P: launchpad.Hit{X: 0, Y: 0}, Width: 8, Height: 8}
r.Paint(pad, widgets.ColorGreenFull, 10*time.Millisecond)

fmt.Println("Then clear rectangle with 15ms transition between each pixel")
r.Clear(pad, 15*time.Millisecond)

r2 := widgets.Rectangle{P: launchpad.Hit{X: 0, Y: 0}, Width: 2, Height: 2}
r2.Paint(pad, widgets.ColorGreenFull, 10*time.Millisecond)
r2.ScrollTo(pad, launchpad.Hit{X: 6, Y: 6}, widgets.ColorGreenFull, 300*time.Millisecond)
```

## Links

* Rakyll Launchpad Lib: https://github.com/rakyll/launchpad
* Launchpad Programmer's Reference Guide: https://global.novationmusic.com/support/downloads/launchpad-programmers-reference-guide

## License

3-clause BSD
