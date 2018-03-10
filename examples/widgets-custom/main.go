package main

import (
	"fmt"
	"time"

	"github.com/rakyll/launchpad"

	"github.com/yesnault/launchpadx"
)

func main() {
	pad, err := launchpad.Open()
	if err != nil {
		fmt.Printf("error while opening connection to launchpad: %v", err)
	}
	defer pad.Close()
	pad.Clear()

	c := padx.NewCustom([]string{
		"--0--",
		"--0--",
		"00000",
		"--0--",
		"-0-0-",
		"0---0",
	})

	c.Paint(pad, padx.ColorRedFull, 0)

	time.Sleep(1 * time.Second)

	c.ScrollTo(pad, 2, 1, padx.ColorGreenFull, 100*time.Millisecond)
	time.Sleep(1 * time.Second)
}
