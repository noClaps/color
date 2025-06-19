package color

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type RGBA struct {
	R float64
	G float64
	B float64
	A float64
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

	rStr := strings.TrimSpace(splits[0])
	gStr := strings.TrimSpace(splits[1])
	bStr := strings.TrimSpace(splits[2])
	aStr := "1"
	if strings.Contains(bStr, "/") {
		splits := strings.SplitN(bStr, "/", 2)
		bStr, aStr = strings.TrimSpace(splits[0]), strings.TrimSpace(splits[1])
	}

	var r float64
	if strings.Contains(rStr, ".") {
		rFloat, err := strconv.ParseFloat(rStr, 64)
		if err != nil {
			return RGBA{}, fmt.Errorf("Error parsing RGB red: %v", err)
		}
		if rFloat < 0 || rFloat > 1 {
			return RGBA{}, fmt.Errorf("Red must be in range 0 ≤ r ≤ 1: `%v`", rFloat)
		}
		r = rFloat
	} else {
		rInt, err := strconv.ParseInt(rStr, 10, 64)
		if err != nil {
			return RGBA{}, fmt.Errorf("Error parsing RGB red: %v", err)
		}
		if rInt < 0 || rInt > 255 {
			return RGBA{}, fmt.Errorf("Red must be in range 0 ≤ r ≤ 255: `%v`", rInt)
		}
		r = float64(rInt) / 255
	}

	var g float64
	if strings.Contains(gStr, ".") {
		gFloat, err := strconv.ParseFloat(gStr, 64)
		if err != nil {
			return RGBA{}, fmt.Errorf("Error parsing RGB green: %v", err)
		}
		if gFloat < 0 || gFloat > 1 {
			return RGBA{}, fmt.Errorf("Green must be in range 0 ≤ g ≤ 1: `%v`", gFloat)
		}
		g = gFloat
	} else {
		gInt, err := strconv.ParseInt(gStr, 10, 64)
		if err != nil {
			return RGBA{}, fmt.Errorf("Error parsing RGB green: %v", err)
		}
		if gInt < 0 || gInt > 255 {
			return RGBA{}, fmt.Errorf("Green must be in range 0 ≤ g ≤ 255: `%v`", gInt)
		}
		g = float64(gInt) / 255
	}

	var b float64
	if strings.Contains(bStr, ".") {
		bFloat, err := strconv.ParseFloat(bStr, 64)
		if err != nil {
			return RGBA{}, fmt.Errorf("Error parsing RGB blue: %v", err)
		}
		if bFloat < 0 || bFloat > 1 {
			return RGBA{}, fmt.Errorf("Blue must be in range 0 ≤ b ≤ 1: `%v`", bFloat)
		}
		b = bFloat
	} else {
		bInt, err := strconv.ParseInt(bStr, 10, 64)
		if err != nil {
			return RGBA{}, fmt.Errorf("Error parsing RGB blue: %v", err)
		}
		if bInt < 0 || bInt > 255 {
			return RGBA{}, fmt.Errorf("Blue must be in range 0 ≤ b ≤ 255: `%v`", bInt)
		}
		b = float64(bInt) / 255
	}

	var a float64
	if aStr[len(aStr)-1] == '%' {
		aFloat, err := strconv.ParseFloat(aStr[:len(aStr)-1], 64)
		if err != nil {
			return RGBA{}, fmt.Errorf("Error parsing RGB alpha: %v", err)
		}
		if aFloat < 0 || aFloat > 100 {
			return RGBA{}, fmt.Errorf("Alpha must be in range 0%% ≤ a ≤ 100%%: `%v`", aFloat)
		}
		a = aFloat / 100
	} else {
		aFloat, err := strconv.ParseFloat(aStr, 64)
		if err != nil {
			return RGBA{}, fmt.Errorf("Error parsing RGB alpha: %v", err)
		}
		if aFloat < 0 || aFloat > 1 {
			return RGBA{}, fmt.Errorf("Alpha must be in range 0 ≤ a ≤ 1: `%v`", a)
		}
		a = aFloat
	}

	return RGBA{r, g, b, a}, nil
}

func (rgb RGBA) ToOklch() Oklch {
	// Convert from sRGB to linear sRGB
	// https://bottosson.github.io/posts/colorwrong/#what-can-we-do
	r := rgb.R / 12.92
	if rgb.R >= 0.04045 {
		r = math.Pow((rgb.R+0.055)/(1+0.055), 2.4)
	}
	g := rgb.G / 12.92
	if rgb.G >= 0.04045 {
		g = math.Pow((rgb.G+0.055)/(1+0.055), 2.4)
	}
	bl := rgb.B / 12.92
	if rgb.B >= 0.04045 {
		bl = math.Pow((rgb.B+0.055)/(1+0.055), 2.4)
	}

	l := 0.4122214708*r + 0.5363325363*g + 0.0514459929*bl
	m := 0.2119034982*r + 0.6806995451*g + 0.1073969566*bl
	s := 0.0883024619*r + 0.2817188376*g + 0.6299787005*bl

	l = math.Cbrt(l)
	m = math.Cbrt(m)
	s = math.Cbrt(s)

	L := 0.2104542553*l + 0.7936177850*m - 0.0040720468*s
	a := 1.9779984951*l - 2.4285922050*m + 0.4505937099*s
	b := 0.0259040371*l + 0.7827717662*m - 0.8086757660*s

	return Oklch{L, math.Hypot(a, b), math.Atan2(b, a), rgb.A}
}

func (rgb RGBA) ToHex() Hex {
	r := uint8(math.Round(rgb.R * 255))
	g := uint8(math.Round(rgb.G * 255))
	b := uint8(math.Round(rgb.B * 255))
	a := uint8(math.Round(rgb.A * 255))
	return Hex{r, g, b, a}
}

func (rgb RGBA) String() string {
	a := ""
	if rgb.A != 1 {
		a = fmt.Sprintf(" / %v%%", rgb.A*100)
	}
	return fmt.Sprintf("rgb(%v %v %v%s)", math.Round(rgb.R*255), math.Round(rgb.G*255), math.Round(rgb.B*255), a)
}
