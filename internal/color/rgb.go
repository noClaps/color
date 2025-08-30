package color

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type RGBA struct {
	Red, Green, Blue, Alpha uint8
}

func NewRGBA(input string) (RGBA, error) {
	input = input[4 : len(input)-1] // remove `rgb(` and `)`

	splits := []string{}
	if strings.Contains(input, ",") {
		splits = strings.SplitN(input, ",", 3)
	} else {
		splits = strings.SplitN(input, " ", 3)
	}

	if len(splits) < 3 {
		return RGBA{}, fmt.Errorf("Invalid RGB input: `rgb(%v)`", input)
	}

	redStr := strings.TrimSpace(splits[0])
	greenStr := strings.TrimSpace(splits[1])
	blueStr := strings.TrimSpace(splits[2])
	alphaStr := "1"
	if strings.Contains(blueStr, "/") {
		splits := strings.SplitN(blueStr, "/", 2)
		blueStr, alphaStr = strings.TrimSpace(splits[0]), strings.TrimSpace(splits[1])
	}

	var red uint8
	if strings.Contains(redStr, ".") {
		redFloat, err := strconv.ParseFloat(redStr, 64)
		if err != nil {
			return RGBA{}, fmt.Errorf("Error parsing RGB red: %v", err)
		}
		if redFloat < 0 || redFloat > 1 {
			return RGBA{}, fmt.Errorf("Red must be in range 0 ≤ r ≤ 1: `%v`", redFloat)
		}
		red = uint8(math.Round(redFloat * 255))
	} else {
		redInt, err := strconv.ParseUint(redStr, 10, 8)
		if err != nil {
			return RGBA{}, fmt.Errorf("Error parsing RGB red: %v", err)
		}
		if redInt > 255 {
			return RGBA{}, fmt.Errorf("Red must be in range 0 ≤ r ≤ 255: `%v`", redInt)
		}
		red = uint8(redInt)
	}

	var green uint8
	if strings.Contains(greenStr, ".") {
		greenFloat, err := strconv.ParseFloat(greenStr, 64)
		if err != nil {
			return RGBA{}, fmt.Errorf("Error parsing RGB green: %v", err)
		}
		if greenFloat < 0 || greenFloat > 1 {
			return RGBA{}, fmt.Errorf("Green must be in range 0 ≤ g ≤ 1: `%v`", greenFloat)
		}
		green = uint8(math.Round(greenFloat * 255))
	} else {
		greenInt, err := strconv.ParseUint(greenStr, 10, 8)
		if err != nil {
			return RGBA{}, fmt.Errorf("Error parsing RGB green: %v", err)
		}
		if greenInt > 255 {
			return RGBA{}, fmt.Errorf("Green must be in range 0 ≤ g ≤ 255: `%v`", greenInt)
		}
		green = uint8(greenInt)
	}

	var blue uint8
	if strings.Contains(blueStr, ".") {
		blueFloat, err := strconv.ParseFloat(blueStr, 64)
		if err != nil {
			return RGBA{}, fmt.Errorf("Error parsing RGB blue: %v", err)
		}
		if blueFloat < 0 || blueFloat > 1 {
			return RGBA{}, fmt.Errorf("Blue must be in range 0 ≤ b ≤ 1: `%v`", blueFloat)
		}
		blue = uint8(math.Round(blueFloat * 255))
	} else {
		blueInt, err := strconv.ParseUint(blueStr, 10, 8)
		if err != nil {
			return RGBA{}, fmt.Errorf("Error parsing RGB blue: %v", err)
		}
		if blueInt > 255 {
			return RGBA{}, fmt.Errorf("Blue must be in range 0 ≤ b ≤ 255: `%v`", blueInt)
		}
		blue = uint8(blueInt)
	}

	var alpha uint8
	if alphaStr[len(alphaStr)-1] == '%' {
		alphaFloat, err := strconv.ParseFloat(alphaStr[:len(alphaStr)-1], 64)
		if err != nil {
			return RGBA{}, fmt.Errorf("Error parsing RGB alpha: %v", err)
		}
		if alphaFloat < 0 || alphaFloat > 100 {
			return RGBA{}, fmt.Errorf("Alpha must be in range 0%% ≤ a ≤ 100%%: `%v`", alphaFloat)
		}
		alpha = uint8(math.Round(alphaFloat / 100 * 255))
	} else {
		alphaFloat, err := strconv.ParseFloat(alphaStr, 64)
		if err != nil {
			return RGBA{}, fmt.Errorf("Error parsing RGB alpha: %v", err)
		}
		if alphaFloat < 0 || alphaFloat > 1 {
			return RGBA{}, fmt.Errorf("Alpha must be in range 0 ≤ a ≤ 1: `%v`", alpha)
		}
		alpha = uint8(math.Round(alphaFloat * 255))
	}

	return RGBA{red, green, blue, alpha}, nil
}

func (rgb RGBA) ToOklch() Oklch {
	red := float64(rgb.Red) / 255
	green := float64(rgb.Green) / 255
	blue := float64(rgb.Blue) / 255
	// Convert from sRGB to linear sRGB
	// https://bottosson.github.io/posts/colorwrong/#what-can-we-do
	linearRed := red / 12.92
	if red >= 0.04045 {
		linearRed = math.Pow((red+0.055)/(1+0.055), 2.4)
	}
	linearGreen := green / 12.92
	if green >= 0.04045 {
		linearGreen = math.Pow((green+0.055)/(1+0.055), 2.4)
	}
	linearBlue := blue / 12.92
	if blue >= 0.04045 {
		linearBlue = math.Pow((blue+0.055)/(1+0.055), 2.4)
	}

	l := 0.4122214708*linearRed + 0.5363325363*linearGreen + 0.0514459929*linearBlue
	m := 0.2119034982*linearRed + 0.6806995451*linearGreen + 0.1073969566*linearBlue
	s := 0.0883024619*linearRed + 0.2817188376*linearGreen + 0.6299787005*linearBlue

	l = math.Cbrt(l)
	m = math.Cbrt(m)
	s = math.Cbrt(s)

	L := 0.2104542553*l + 0.7936177850*m - 0.0040720468*s
	a := 1.9779984951*l - 2.4285922050*m + 0.4505937099*s
	b := 0.0259040371*l + 0.7827717662*m - 0.8086757660*s

	L = closeEnough(L)
	a = closeEnough(a)
	b = closeEnough(b)

	hue := math.Atan2(b, a) / math.Pi * 180 // rad -> deg
	// Bring hue between 0 <= h <= 360
	for hue > 360 {
		hue -= 360
	}
	for hue < 0 {
		hue += 360
	}
	return Oklch{L, math.Hypot(a, b), hue, float64(rgb.Alpha) / 255}
}

func (rgb RGBA) ToHex() Hex {
	return Hex{rgb.Red, rgb.Green, rgb.Blue, rgb.Alpha}
}

func (rgb RGBA) String() string {
	alpha := ""
	if rgb.Alpha != 255 {
		alpha = fmt.Sprintf(" / %v%%", float64(rgb.Alpha)/255*100)
	}
	return fmt.Sprintf("rgb(%v %v %v%s)", rgb.Red, rgb.Green, rgb.Blue, alpha)
}

func closeEnough(f float64) float64 {
	if math.Abs(math.Round(f)-f) < 1e-5 {
		return math.Round(f)
	}
	return f
}
