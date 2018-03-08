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
	"github.com/yesnault/launchpadx/fonts"
)


func main() {
	pad, err := launchpad.Open()
	if err != nil {
		fmt.Printf("error while openning connection to launchpad: %v", err)
	}
	defer pad.Close()

	fmt.Println("Press a button a-h (vertical) or 1-8 (horizontal) on the launchpad")

	pad.Clear()
	f := fonts.Size8x8()
	ch := pad.Listen()
	for {
		hit := <-ch
		pad.Clear()

		// check if the key is a menu button : a-h, 1-8
		if btn := buttons.Get(hit); btn != nil {
			// it's a 'button'

			// get a random color
			color := fonts.Colors[rand.Intn(len(fonts.Colors))]

			// display the key pressed
			fmt.Println("pressed:", hit, "with color:", color.Name)

			f.CharSet[btn.Name].Paint(pad, color)
		} else {
			fmt.Println("pressed (not a button):", hit)
		}
	}
}

```

### Text

```go
[...]

text := fonts.NewText("foobar", fonts.DirectionLeftToRight, fonts.Size8x8())
text.Scroll(pad, fonts.RedFull, fonts.DirectionRightToLeft, 100*time.Millisecond)

[...]
```


## Manipulate a character

```go
// initialize the font
f := fonts.Size8x8()

// paint a character
f.CharSet["a"].Paint(pad, fonts.GreenFull)

// blink a character, 100ms transition, 15 repeats between two colors
f.CharSet["b"].Blink(pad, fonts.RedFull, fonts.Off, 100*time.Millisecond, 15)

// Scroll a character, 500ms transition
f.CharSet["c"].Scroll(pad, fonts.RedFull, fonts.DirectionTopToBottom, 500*time.Millisecond)

```

## TODO

- [x] Blink
- [x] Scrolling
- [ ] Font 8x8 more symbols
- [ ] Font 4x4
- [x] Text Display

## Links

* Rakyll Launchpad Lib: https://github.com/rakyll/launchpad
* Launchpad Programmer's Reference Guide: https://global.novationmusic.com/support/downloads/launchpad-programmers-reference-guide

## License

3-clause BSD
