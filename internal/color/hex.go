package color

import (
	"fmt"
	"strconv"
)

type Hex struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func NewHex(input string) (Hex, error) {
	input = input[1:] // skip #
	if len(input) == 3 || len(input) == 4 {
		newInput := ""
		for _, char := range input {
			newInput += fmt.Sprintf("%c%c", char, char)
		}
		input = newInput
	}
	if len(input) != 6 && len(input) != 8 {
		return Hex{}, fmt.Errorf("Invalid hex code: `%v`", input)
	}

	r, err := strconv.ParseUint(string(input[0:2]), 16, 8)
	if err != nil {
		return Hex{}, fmt.Errorf("Error parsing hex: %v", err)
	}
	g, err := strconv.ParseUint(string(input[2:4]), 16, 8)
	if err != nil {
		return Hex{}, fmt.Errorf("Error parsing hex: %v", err)
	}
	b, err := strconv.ParseUint(string(input[4:6]), 16, 8)
	if err != nil {
		return Hex{}, fmt.Errorf("Error parsing hex: %v", err)
	}

	var a uint8 = 255
	if len(input) > 6 {
		alpha, err := strconv.ParseUint(string(input[6:8]), 16, 8)
		if err != nil {
			return Hex{}, fmt.Errorf("Error parsing hex: %v", err)
		}
		a = uint8(alpha)
	}

	return Hex{uint8(r), uint8(g), uint8(b), a}, nil
}

func (h Hex) ToOklch() Oklch {
	return h.ToRGBA().ToOklch()
}

func (h Hex) ToRGBA() RGBA {
	return RGBA{float64(h.R) / 255, float64(h.G) / 255, float64(h.B) / 255, float64(h.A) / 255}
}

func (h Hex) String() string {
	if h.A == 255 {
		return fmt.Sprintf("#%02x%02x%02x", h.R, h.G, h.B)
	}
	return fmt.Sprintf("#%02x%02x%02x%02x", h.R, h.G, h.B, h.A)
}
