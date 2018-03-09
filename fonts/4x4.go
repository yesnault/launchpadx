package fonts

// Size4x4 return a 8 per 8 pixels font
func Size4x4() Font {
	return Font{
		Size: Size{Height: 4, Width: 4},
		CharSet: map[rune]Character{
			'a': Character{
				Hits: process([]string{
					"-0-",
					"0-0",
					"000",
					"0-0",
				}),
			},
			'b': Character{
				Hits: process([]string{
					"0--",
					"000",
					"0-0",
					"000",
				}),
			},
			'c': Character{
				Hits: process([]string{
					"000",
					"0--",
					"0--",
					"000",
				}),
			},
			'd': Character{
				Hits: process([]string{
					"--0",
					"000",
					"0-0",
					"000",
				}),
			},
			'e': Character{
				Hits: process([]string{
					"000",
					"00-",
					"0-",
					"000",
				}),
			},
			'f': Character{
				Hits: process([]string{
					"000",
					"00-",
					"0-",
					"0--",
				}),
			},
			'g': Character{
				Hits: process([]string{
					"000",
					"0--",
					"000",
					"000",
				}),
			},
			'h': Character{
				Hits: process([]string{
					"0-0",
					"0-0",
					"000",
					"0-0",
				}),
			},
			'i': Character{
				Hits: process([]string{
					"000",
					"-0-",
					"-0-",
					"000",
				}),
			},
			'j': Character{
				Hits: process([]string{
					"-00",
					"--0",
					"-00",
					"000",
				}),
			},
			'k': Character{
				Hits: process([]string{
					"0-0",
					"00-",
					"00-",
					"0-0",
				}),
			},
			'l': Character{
				Hits: process([]string{
					"0--",
					"0--",
					"0--",
					"000",
				}),
			},
			'm': Character{
				Hits: process([]string{
					"0-0",
					"000",
					"0-0",
					"0-0",
				}),
			},
			'n': Character{
				Hits: process([]string{
					"0--",
					"000",
					"0-0",
					"0-0",
				}),
			},
			'o': Character{
				Hits: process([]string{
					"-0-",
					"0-0",
					"0-0",
					"-0-",
				}),
			},
			'p': Character{
				Hits: process([]string{
					"000",
					"000",
					"0--",
					"0--",
				}),
			},
			'q': Character{
				Hits: process([]string{
					"000",
					"000",
					"--0",
					"--0",
				}),
			},
			'r': Character{
				Hits: process([]string{
					"000",
					"0--",
					"0--",
					"0--",
				}),
			},
			's': Character{
				Hits: process([]string{
					"000",
					"0--",
					"-00",
					"000",
				}),
			},
			't': Character{
				Hits: process([]string{
					"000",
					"-0-",
					"-0-",
					"-0-",
				}),
			},
			'u': Character{
				Hits: process([]string{
					"0-0",
					"0-0",
					"0-0",
					"000",
				}),
			},
			'v': Character{
				Hits: process([]string{
					"0-0",
					"0-0",
					"0-0",
					"-0-",
				}),
			},
			'w': Character{
				Hits: process([]string{
					"0-0",
					"000",
					"000",
					"0-0",
				}),
			},
			'x': Character{
				Hits: process([]string{
					"0-0",
					"-0-",
					"0-0",
					"0-0",
				}),
			},
			'y': Character{
				Hits: process([]string{
					"0-0",
					"-00",
					"--0",
					"--0",
				}),
			},
			'z': Character{
				Hits: process([]string{
					"000",
					"-0-",
					"0--",
					"000",
				}),
			},
			'1': Character{
				Hits: process([]string{
					"-0-",
					"00-",
					"-0-",
					"000",
				}),
			},
			'2': Character{
				Hits: process([]string{
					"000",
					"-00",
					"0--",
					"000",
				}),
			},
			'3': Character{
				Hits: process([]string{
					"000",
					"-00",
					"--0",
					"000",
				}),
			},
			'4': Character{
				Hits: process([]string{
					"0-0",
					"0-0",
					"000",
					"--0",
				}),
			},
			'5': Character{
				Hits: process([]string{
					"000",
					"00-",
					"--0",
					"000",
				}),
			},
			'6': Character{
				Hits: process([]string{
					"-00",
					"0--",
					"000",
					"000",
				}),
			},
			'7': Character{
				Hits: process([]string{
					"000",
					"--0",
					"-0-",
					"-0-",
				}),
			},
			'8': Character{
				Hits: process([]string{
					"000",
					"000",
					"0-0",
					"000",
				}),
			},
			'9': Character{
				Hits: process([]string{
					"000",
					"000",
					"--0",
					"00-",
				}),
			},
			'0': Character{
				Hits: process([]string{
					"000",
					"0-0",
					"0-0",
					"000",
				}),
			},
		},
	}
}
