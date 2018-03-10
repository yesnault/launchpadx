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

	padx "github.com/yesnault/launchpadx"
)


func main() {
	pad, err := launchpad.Open()
	if err != nil {
		fmt.Printf("error while openning connection to launchpad: %v", err)
	}
	defer pad.Close()

	fmt.Println("Press a button a-h (vertical) or 1-8 (horizontal) on the launchpad")

	pad.Clear()
	f := padx.Widgets8x8()
	ch := pad.Listen()
	for {
		hit := <-ch
		pad.Clear()

		// check if the key is a menu button : a-h, 1-8
		if btn := buttons.Get(hit); btn != nil {
			// it's a 'button'

			// get a random color
			color := padx.Colors[rand.Intn(len(padx.Colors))]

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

text := padx.Text("foobar", padx.DirectionLeftToRight, padx.Widgets8x8())
text.Scroll(pad, padx.ColorRedFull, padx.DirectionRightToLeft, 100*time.Millisecond)

[...]
```


## Draw a character

```go
// initialize the font
f := padx.Widgets8x8()

// paint a character
f.Widgets["a"].Paint(pad, padx.ColorGreenFull)

// blink a character, 100ms transition, 15 repeats between two colors
f.Widgets["b"].Blink(pad, padx.ColorRedFull, padx.ColorOff, 100*time.Millisecond, 15)

// Scroll a character, 500ms transition
f.Widgets["c"].Scroll(pad, padx.ColorRedFull, padx.DirectionTopToBottom, 500*time.Millisecond)

```

## Draw a point

```go
padx.Point(5, 3).Paint(pad, padx.ColorRedFull, 0)
```

## Draw a rectangle

```go
r := padx.Rectangle(0, 0, 8, 8)
r.Paint(pad, padx.ColorGreenFull, 10*time.Millisecond)

fmt.Println("Then clear rectangle with 15ms transition between each pixel")
r.Clear(pad, 15*time.Millisecond)

r2 := padx.Rectangle(0, 0, 2, 2)
r2.Paint(pad, padx.ColorGreenFull, 10*time.Millisecond)
r2.ScrollTo(pad, 6, 6, padx.ColorGreenFull, 300*time.Millisecond)
```

## Custom Draw

```go
c := padx.NewCustom([]string{
	"--0--",
	"--0--",
	"00000",
	"--0--",
	"-0-0-",
	"0---0",
})

c.Paint(pad, padx.ColorRedFull, 0)
time.Sleep(100*time.Millisecond)

c.ScrollTo(pad, 2, 1, padx.ColorGreenFull, 100*time.Millisecond)
```

## Links

* Rakyll Launchpad Lib: https://github.com/rakyll/launchpad
* Launchpad Programmer's Reference Guide: https://global.novationmusic.com/support/downloads/launchpad-programmers-reference-guide

## License

3-clause BSD
