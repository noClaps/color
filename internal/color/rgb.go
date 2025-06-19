package color

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type RGBA struct {
	Red   float64
	Green float64
	Blue  float64
	Alpha float64
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

	var red float64
	if strings.Contains(redStr, ".") {
		redFloat, err := strconv.ParseFloat(redStr, 64)
		if err != nil {
			return RGBA{}, fmt.Errorf("Error parsing RGB red: %v", err)
		}
		if redFloat < 0 || redFloat > 1 {
			return RGBA{}, fmt.Errorf("Red must be in range 0 ≤ r ≤ 1: `%v`", redFloat)
		}
		red = redFloat
	} else {
		redInt, err := strconv.ParseInt(redStr, 10, 64)
		if err != nil {
			return RGBA{}, fmt.Errorf("Error parsing RGB red: %v", err)
		}
		if redInt < 0 || redInt > 255 {
			return RGBA{}, fmt.Errorf("Red must be in range 0 ≤ r ≤ 255: `%v`", redInt)
		}
		red = float64(redInt) / 255
	}

	var green float64
	if strings.Contains(greenStr, ".") {
		greenFloat, err := strconv.ParseFloat(greenStr, 64)
		if err != nil {
			return RGBA{}, fmt.Errorf("Error parsing RGB green: %v", err)
		}
		if greenFloat < 0 || greenFloat > 1 {
			return RGBA{}, fmt.Errorf("Green must be in range 0 ≤ g ≤ 1: `%v`", greenFloat)
		}
		green = greenFloat
	} else {
		greenInt, err := strconv.ParseInt(greenStr, 10, 64)
		if err != nil {
			return RGBA{}, fmt.Errorf("Error parsing RGB green: %v", err)
		}
		if greenInt < 0 || greenInt > 255 {
			return RGBA{}, fmt.Errorf("Green must be in range 0 ≤ g ≤ 255: `%v`", greenInt)
		}
		green = float64(greenInt) / 255
	}

	var blue float64
	if strings.Contains(blueStr, ".") {
		blueFloat, err := strconv.ParseFloat(blueStr, 64)
		if err != nil {
			return RGBA{}, fmt.Errorf("Error parsing RGB blue: %v", err)
		}
		if blueFloat < 0 || blueFloat > 1 {
			return RGBA{}, fmt.Errorf("Blue must be in range 0 ≤ b ≤ 1: `%v`", blueFloat)
		}
		blue = blueFloat
	} else {
		blueInt, err := strconv.ParseInt(blueStr, 10, 64)
		if err != nil {
			return RGBA{}, fmt.Errorf("Error parsing RGB blue: %v", err)
		}
		if blueInt < 0 || blueInt > 255 {
			return RGBA{}, fmt.Errorf("Blue must be in range 0 ≤ b ≤ 255: `%v`", blueInt)
		}
		blue = float64(blueInt) / 255
	}

	var alpha float64
	if alphaStr[len(alphaStr)-1] == '%' {
		alphaFloat, err := strconv.ParseFloat(alphaStr[:len(alphaStr)-1], 64)
		if err != nil {
			return RGBA{}, fmt.Errorf("Error parsing RGB alpha: %v", err)
		}
		if alphaFloat < 0 || alphaFloat > 100 {
			return RGBA{}, fmt.Errorf("Alpha must be in range 0%% ≤ a ≤ 100%%: `%v`", alphaFloat)
		}
		alpha = alphaFloat / 100
	} else {
		alphaFloat, err := strconv.ParseFloat(alphaStr, 64)
		if err != nil {
			return RGBA{}, fmt.Errorf("Error parsing RGB alpha: %v", err)
		}
		if alphaFloat < 0 || alphaFloat > 1 {
			return RGBA{}, fmt.Errorf("Alpha must be in range 0 ≤ a ≤ 1: `%v`", alpha)
		}
		alpha = alphaFloat
	}

	return RGBA{red, green, blue, alpha}, nil
}

func (rgb RGBA) ToOklch() Oklch {
	// Convert from sRGB to linear sRGB
	// https://bottosson.github.io/posts/colorwrong/#what-can-we-do
	red := rgb.Red / 12.92
	if rgb.Red >= 0.04045 {
		red = math.Pow((rgb.Red+0.055)/(1+0.055), 2.4)
	}
	green := rgb.Green / 12.92
	if rgb.Green >= 0.04045 {
		green = math.Pow((rgb.Green+0.055)/(1+0.055), 2.4)
	}
	blue := rgb.Blue / 12.92
	if rgb.Blue >= 0.04045 {
		blue = math.Pow((rgb.Blue+0.055)/(1+0.055), 2.4)
	}

	l := 0.4122214708*red + 0.5363325363*green + 0.0514459929*blue
	m := 0.2119034982*red + 0.6806995451*green + 0.1073969566*blue
	s := 0.0883024619*red + 0.2817188376*green + 0.6299787005*blue

	l = math.Cbrt(l)
	m = math.Cbrt(m)
	s = math.Cbrt(s)

	L := 0.2104542553*l + 0.7936177850*m - 0.0040720468*s
	a := 1.9779984951*l - 2.4285922050*m + 0.4505937099*s
	b := 0.0259040371*l + 0.7827717662*m - 0.8086757660*s

	return Oklch{L, math.Hypot(a, b), math.Atan2(b, a), rgb.Alpha}
}

func (rgb RGBA) ToHex() Hex {
	red := uint8(math.Round(rgb.Red * 255))
	green := uint8(math.Round(rgb.Green * 255))
	blue := uint8(math.Round(rgb.Blue * 255))
	alpha := uint8(math.Round(rgb.Alpha * 255))
	return Hex{red, green, blue, alpha}
}

func (rgb RGBA) String() string {
	alpha := ""
	if rgb.Alpha != 1 {
		alpha = fmt.Sprintf(" / %v%%", rgb.Alpha*100)
	}
	return fmt.Sprintf("rgb(%v %v %v%s)", math.Round(rgb.Red*255), math.Round(rgb.Green*255), math.Round(rgb.Blue*255), alpha)
}
