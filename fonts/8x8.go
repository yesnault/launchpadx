package fonts

// Size8x8 return a 8 per 8 pixels font
func Size8x8() Font {
	return Font{
		Size: Size{Height: 8, Width: 8},
		CharSet: map[string]Character{
			"a": Character{
				Hits: process([]string{
					"--------",
					"---000--",
					"--0---0-",
					"--0---0-",
					"--00000-",
					"--0---0-",
					"--0---0-",
					"--0---0-",
				}),
			},
			"b": Character{
				Hits: process([]string{
					"--------",
					"--0000--",
					"--0---0-",
					"--0---0-",
					"--0000--",
					"--0---0-",
					"--0---0-",
					"--00000-",
				}),
			},
			"c": Character{
				Hits: process([]string{
					"--------",
					"---000--",
					"--0---0-",
					"--0-----",
					"--0-----",
					"--0-----",
					"--0---0-",
					"---000--",
				}),
			},
			"d": Character{
				Hits: process([]string{
					"--------",
					"--0000--",
					"--0---0-",
					"--0---0-",
					"--0---0-",
					"--0---0-",
					"--0---0-",
					"--0000--",
				}),
			},
			"e": Character{
				Hits: process([]string{
					"--------",
					"--00000-",
					"--0-----",
					"--0-----",
					"--0000--",
					"--0-----",
					"--0-----",
					"--00000-",
				}),
			},
			"f": Character{
				Hits: process([]string{
					"--------",
					"--00000-",
					"--0-----",
					"--0-----",
					"--0000--",
					"--0-----",
					"--0-----",
					"--0-----",
				}),
			},
			"g": Character{
				Hits: process([]string{
					"--------",
					"---000--",
					"--0---0-",
					"--0-----",
					"--0-----",
					"--0--00-",
					"--0---0-",
					"---000--",
				}),
			},
			"h": Character{
				Hits: process([]string{
					"--------",
					"--0---0-",
					"--0---0-",
					"--0---0-",
					"--00000-",
					"--0---0-",
					"--0---0-",
					"--0---0-",
				}),
			},
			"i": Character{
				Hits: process([]string{
					"--------",
					"---000--",
					"----0---",
					"----0---",
					"----0---",
					"----0---",
					"----0---",
					"---000--",
				}),
			},
			"j": Character{
				Hits: process([]string{
					"--------",
					"----000-",
					"-----0--",
					"-----0--",
					"-----0--",
					"-----0--",
					"--0--0--",
					"---00---",
				}),
			},
			"k": Character{
				Hits: process([]string{
					"--------",
					"--0---0-",
					"--0--0--",
					"--0-0---",
					"--00----",
					"--0-0---",
					"--0--0--",
					"--0---0-",
				}),
			},
			"l": Character{
				Hits: process([]string{
					"--------",
					"--0-----",
					"--0-----",
					"--0-----",
					"--0-----",
					"--0-----",
					"--0-----",
					"--00000-",
				}),
			},
			"m": Character{
				Hits: process([]string{
					"--------",
					"--0---0-",
					"--00-00-",
					"--0-0-0-",
					"--0-0-0-",
					"--0---0-",
					"--0---0-",
					"--0---0-",
				}),
			},
			"n": Character{
				Hits: process([]string{
					"--------",
					"--0---0-",
					"--0---0-",
					"--00--0-",
					"--0-0-0-",
					"--0--00-",
					"--0---0-",
					"--0---0-",
				}),
			},
			"o": Character{
				Hits: process([]string{
					"--------",
					"---000--",
					"--0---0-",
					"--0---0-",
					"--0---0-",
					"--0---0-",
					"--0---0-",
					"---000--",
				}),
			},
			"p": Character{
				Hits: process([]string{
					"--------",
					"--0000--",
					"--0---0-",
					"--0---0-",
					"--0000--",
					"--0-----",
					"--0-----",
					"--0-----",
				}),
			},
			"q": Character{
				Hits: process([]string{
					"--------",
					"---000--",
					"--0---0-",
					"--0---0-",
					"--0---0-",
					"--0-0-0-",
					"--0--0--",
					"---00-0-",
				}),
			},
			"r": Character{
				Hits: process([]string{
					"--------",
					"--0000--",
					"--0---0-",
					"--0---0-",
					"--0000--",
					"--0-0---",
					"--0--0--",
					"--0---0-",
				}),
			},
			"s": Character{
				Hits: process([]string{
					"--------",
					"---0000-",
					"--0-----",
					"--0-----",
					"---000--",
					"------0-",
					"------0-",
					"--0000--",
				}),
			},
			"t": Character{
				Hits: process([]string{
					"--------",
					"--00000-",
					"----0---",
					"----0---",
					"----0---",
					"----0---",
					"----0---",
					"----0---",
				}),
			},
			"u": Character{
				Hits: process([]string{
					"--------",
					"--0---0-",
					"--0---0-",
					"--0---0-",
					"--0---0-",
					"--0---0-",
					"--0---0-",
					"---000--",
				}),
			},
			"v": Character{
				Hits: process([]string{
					"--------",
					"--0---0-",
					"--0---0-",
					"--0---0-",
					"--0---0-",
					"---0-0--",
					"---0-0--",
					"----0---",
				}),
			},
			"w": Character{
				Hits: process([]string{
					"--------",
					"--0---0-",
					"--0---0-",
					"--0---0-",
					"--0-0-0-",
					"--0-0-0-",
					"--0-0-0-",
					"---0-0--",
				}),
			},
			"x": Character{
				Hits: process([]string{
					"--------",
					"--0---0-",
					"--0---0-",
					"---0-0--",
					"----0---",
					"---0-0--",
					"--0---0-",
					"--0---0-",
				}),
			},
			"y": Character{
				Hits: process([]string{
					"--------",
					"--0---0-",
					"--0---0-",
					"---0-0--",
					"----0---",
					"----0---",
					"----0---",
					"----0---",
				}),
			},
			"z": Character{
				Hits: process([]string{
					"--------",
					"--00000-",
					"------0-",
					"-----0--",
					"----0---",
					"---0----",
					"--0-----",
					"--00000-",
				}),
			},
			"1": Character{
				Hits: process([]string{
					"--------",
					"----0---",
					"---00---",
					"----0---",
					"----0---",
					"----0---",
					"----0---",
					"---000--",
				}),
			},
			"2": Character{
				Hits: process([]string{
					"--------",
					"---000--",
					"--0---0-",
					"------0-",
					"-----0--",
					"----0---",
					"---0----",
					"--00000-",
				}),
			},
			"3": Character{
				Hits: process([]string{
					"--------",
					"--00000-",
					"-----0--",
					"----0---",
					"-----0--",
					"------0-",
					"--0---0-",
					"---000--",
				}),
			},
			"4": Character{
				Hits: process([]string{
					"--------",
					"-----0--",
					"----00--",
					"---0-0--",
					"--0--0--",
					"--00000-",
					"-----0--",
					"-----0--",
				}),
			},
			"5": Character{
				Hits: process([]string{
					"--------",
					"--00000-",
					"--0-----",
					"--0000--",
					"------0-",
					"------0-",
					"--0---0-",
					"---000--",
				}),
			},
			"6": Character{
				Hits: process([]string{
					"--------",
					"----00--",
					"---0----",
					"--0-----",
					"--0000--",
					"--0---0-",
					"--0---0-",
					"---000--",
				}),
			},
			"7": Character{
				Hits: process([]string{
					"--------",
					"--00000-",
					"------0-",
					"-----0--",
					"----0---",
					"---0----",
					"---0----",
					"---0----",
				}),
			},
			"8": Character{
				Hits: process([]string{
					"--------",
					"---000--",
					"--0---0-",
					"--0---0-",
					"---000--",
					"--0---0-",
					"--0---0-",
					"---000--",
				}),
			},
			"9": Character{
				Hits: process([]string{
					"--------",
					"---000--",
					"--0---0-",
					"--0---0-",
					"---0000-",
					"------0-",
					"-----0--",
					"---00---",
				}),
			},
			"0": Character{
				Hits: process([]string{
					"--------",
					"---000--",
					"--0---0-",
					"--0--00-",
					"--0-0-0-",
					"--00--0-",
					"--0---0-",
					"---000--",
				}),
			},
			"?": Character{
				Hits: process([]string{
					"--------",
					"--000---",
					"-0---0--",
					"----0---",
					"---0----",
					"---0----",
					"--------",
					"---0----",
				}),
			},
			"!": Character{
				Hits: process([]string{
					"--------",
					"----0---",
					"----0---",
					"----0---",
					"----0---",
					"----0---",
					"--------",
					"----0---",
				}),
			},
			"|": Character{
				Hits: process([]string{
					"----0---",
					"----0---",
					"----0---",
					"----0---",
					"----0---",
					"----0---",
					"----0---",
					"----0---",
				}),
			},
			"all": Character{
				Hits: process([]string{
					"00000000",
					"00000000",
					"00000000",
					"00000000",
					"00000000",
					"00000000",
					"00000000",
					"00000000",
				}),
			},
		},
	}
}