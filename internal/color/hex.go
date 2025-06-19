package color

import (
	"fmt"
	"strconv"
)

type Hex struct {
	Red   uint8
	Green uint8
	Blue  uint8
	Alpha uint8
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

	red, err := strconv.ParseUint(string(input[0:2]), 16, 8)
	if err != nil {
		return Hex{}, fmt.Errorf("Error parsing hex: %v", err)
	}
	green, err := strconv.ParseUint(string(input[2:4]), 16, 8)
	if err != nil {
		return Hex{}, fmt.Errorf("Error parsing hex: %v", err)
	}
	blue, err := strconv.ParseUint(string(input[4:6]), 16, 8)
	if err != nil {
		return Hex{}, fmt.Errorf("Error parsing hex: %v", err)
	}

	var alpha uint8 = 255
	if len(input) > 6 {
		alphaUint, err := strconv.ParseUint(string(input[6:8]), 16, 8)
		if err != nil {
			return Hex{}, fmt.Errorf("Error parsing hex: %v", err)
		}
		alpha = uint8(alphaUint)
	}

	return Hex{uint8(red), uint8(green), uint8(blue), alpha}, nil
}

func (h Hex) ToOklch() Oklch {
	return h.ToRGBA().ToOklch()
}

func (h Hex) ToRGBA() RGBA {
	return RGBA{float64(h.Red) / 255, float64(h.Green) / 255, float64(h.Blue) / 255, float64(h.Alpha) / 255}
}

func (h Hex) String() string {
	if h.Alpha == 255 {
		return fmt.Sprintf("#%02x%02x%02x", h.Red, h.Green, h.Blue)
	}
	return fmt.Sprintf("#%02x%02x%02x%02x", h.Red, h.Green, h.Blue, h.Alpha)
}
