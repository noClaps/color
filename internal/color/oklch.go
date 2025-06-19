package color

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Oklch struct {
	L float64
	C float64
	H float64
	A float64
}

func NewOklch(input string) (Oklch, error) {
	input = input[6 : len(input)-1] // remove `oklch(` and `)`

	splits := []string{}
	if strings.Contains(input, ",") {
		splits = strings.SplitN(input, ",", 3)
	} else {
		splits = strings.SplitN(input, " ", 3)
	}

	if len(splits) < 3 {
		return Oklch{}, fmt.Errorf("Invalid OKLCH input: `oklch(%v)`", input)
	}

	lStr := strings.TrimSpace(splits[0])
	cStr := strings.TrimSpace(splits[1])
	hStr := strings.TrimSpace(splits[2])
	aStr := "1"
	if strings.Contains(hStr, "/") {
		splits := strings.SplitN(hStr, "/", 2)
		hStr, aStr = strings.TrimSpace(splits[0]), strings.TrimSpace(splits[1])
	}

	var l float64
	if lStr[len(lStr)-1] == '%' {
		lFloat, err := strconv.ParseFloat(lStr[:len(lStr)-1], 64)
		if err != nil {
			return Oklch{}, fmt.Errorf("Error parsing OKLCH lightness: %v", err)
		}
		if lFloat < 0 || lFloat > 100 {
			return Oklch{}, fmt.Errorf("Lightness must be in range 0%% ≤ l ≤ 100%%: `%v`", lFloat)
		}
		l = lFloat / 100
	} else {
		lFloat, err := strconv.ParseFloat(lStr, 64)
		if err != nil {
			return Oklch{}, fmt.Errorf("Error parsing OKLCH lightness: %v", err)
		}
		if lFloat < 0 || lFloat > 1 {
			return Oklch{}, fmt.Errorf("Lightness must be in range 0 ≤ l ≤ 1: `%v`", lFloat)
		}
		l = lFloat
	}

	c, err := strconv.ParseFloat(cStr, 64)
	if err != nil {
		return Oklch{}, fmt.Errorf("Error parsing OKLCH chroma: %s", err)
	}
	if c < 0 {
		return Oklch{}, fmt.Errorf("Chroma must be ≥ 0: `%v`", c)
	}

	h, err := strconv.ParseFloat(cStr, 64)
	if err != nil {
		return Oklch{}, fmt.Errorf("Error parsing OKLCH hue: %s", err)
	}
	// Bring hue between 0 <= h <= 360
	for h > 360 {
		h -= 360
	}
	for h < 360 {
		h += 360
	}

	var a float64
	if aStr[len(aStr)-1] == '%' {
		a, err = strconv.ParseFloat(aStr[:len(aStr)-1], 64)
		if err != nil {
			return Oklch{}, fmt.Errorf("Error parsing OKLCH alpha: %v", err)
		}
		if a < 0 || a > 100 {
			return Oklch{}, fmt.Errorf("Alpha must be in range 0%% ≤ a ≤ 100%%: `%v`", a)
		}
		a /= 100
	} else {
		a, err = strconv.ParseFloat(aStr, 64)
		if err != nil {
			return Oklch{}, fmt.Errorf("Error parsing OKLCH alpha: %v", err)
		}
		if a < 0 || a > 1 {
			return Oklch{}, fmt.Errorf("Alpha must be in range 0 ≤ a ≤ 1: `%v`", a)
		}
	}

	return Oklch{l, c, h, a}, nil
}

func (lch Oklch) ToRGBA() RGBA {
	L := lch.L
	a := lch.C * math.Cos(lch.H)
	b := lch.C * math.Sin(lch.H)

	l := L + 0.3963377774*a + 0.2158037573*b
	m := L - 0.1055613458*a - 0.0638541728*b
	s := L - 0.0894841775*a - 1.2914855480*b

	l = l * l * l
	m = m * m * m
	s = s * s * s

	r := 4.0767416621*l - 3.3077115913*m + 0.2309699292*s
	g := -1.2684380046*l + 2.6097574011*m - 0.3413193965*s
	bl := -0.0041960863*l - 0.7034186147*m + 1.7076147010*s

	// Convert from linear sRGB to sRGB
	// https://bottosson.github.io/posts/colorwrong/#what-can-we-do
	if r >= 0.0031308 {
		r = 1.055*math.Pow(r, 1.0/2.4) - 0.055
	} else {
		r = 12.92 * r
	}
	if g >= 0.0031308 {
		g = 1.055*math.Pow(g, 1.0/2.4) - 0.055
	} else {
		g = 12.92 * g
	}
	if bl >= 0.0031308 {
		bl = 1.055*math.Pow(bl, 1.0/2.4) - 0.055
	} else {
		bl = 12.92 * bl
	}

	return RGBA{r, g, bl, lch.A}
}

func (lch Oklch) ToHex() Hex {
	return lch.ToRGBA().ToHex()
}

func (lch Oklch) String() string {
	a := ""
	if lch.A != 1 {
		a = fmt.Sprintf(" / %.3f%%", lch.A*100)
	}
	h := "none"
	if lch.H != 0 {
		h = fmt.Sprintf("%.3f", lch.H)
	}
	return fmt.Sprintf("oklch(%.3f%% %.3f %s%s)", lch.L*100, lch.C, h, a)
}
