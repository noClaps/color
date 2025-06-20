package color

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Oklch struct {
	Lightness float64
	Chroma    float64
	Hue       float64
	Alpha     float64
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

	lightnessStr := strings.TrimSpace(splits[0])
	chromaStr := strings.TrimSpace(splits[1])
	hueStr := strings.TrimSpace(splits[2])
	alphaStr := "1"
	if strings.Contains(hueStr, "/") {
		splits := strings.SplitN(hueStr, "/", 2)
		hueStr, alphaStr = strings.TrimSpace(splits[0]), strings.TrimSpace(splits[1])
	}

	var lightness float64
	if lightnessStr[len(lightnessStr)-1] == '%' {
		lightnessFloat, err := strconv.ParseFloat(lightnessStr[:len(lightnessStr)-1], 64)
		if err != nil {
			return Oklch{}, fmt.Errorf("Error parsing OKLCH lightness: %v", err)
		}
		if lightnessFloat < 0 || lightnessFloat > 100 {
			return Oklch{}, fmt.Errorf("Lightness must be in range 0%% ≤ l ≤ 100%%: `%v`", lightnessFloat)
		}
		lightness = lightnessFloat / 100
	} else {
		lightnessFloat, err := strconv.ParseFloat(lightnessStr, 64)
		if err != nil {
			return Oklch{}, fmt.Errorf("Error parsing OKLCH lightness: %v", err)
		}
		if lightnessFloat < 0 || lightnessFloat > 1 {
			return Oklch{}, fmt.Errorf("Lightness must be in range 0 ≤ l ≤ 1: `%v`", lightnessFloat)
		}
		lightness = lightnessFloat
	}

	chroma, err := strconv.ParseFloat(chromaStr, 64)
	if err != nil {
		return Oklch{}, fmt.Errorf("Error parsing OKLCH chroma: %s", err)
	}
	if chroma < 0 {
		return Oklch{}, fmt.Errorf("Chroma must be ≥ 0: `%v`", chroma)
	}

	hue, err := strconv.ParseFloat(chromaStr, 64)
	if err != nil {
		return Oklch{}, fmt.Errorf("Error parsing OKLCH hue: %s", err)
	}
	// Bring hue between 0 <= h <= 360
	for hue > 360 {
		hue -= 360
	}
	for hue < 360 {
		hue += 360
	}

	var alpha float64
	if alphaStr[len(alphaStr)-1] == '%' {
		alpha, err = strconv.ParseFloat(alphaStr[:len(alphaStr)-1], 64)
		if err != nil {
			return Oklch{}, fmt.Errorf("Error parsing OKLCH alpha: %v", err)
		}
		if alpha < 0 || alpha > 100 {
			return Oklch{}, fmt.Errorf("Alpha must be in range 0%% ≤ a ≤ 100%%: `%v`", alpha)
		}
		alpha /= 100
	} else {
		alpha, err = strconv.ParseFloat(alphaStr, 64)
		if err != nil {
			return Oklch{}, fmt.Errorf("Error parsing OKLCH alpha: %v", err)
		}
		if alpha < 0 || alpha > 1 {
			return Oklch{}, fmt.Errorf("Alpha must be in range 0 ≤ a ≤ 1: `%v`", alpha)
		}
	}

	return Oklch{lightness, chroma, hue, alpha}, nil
}

func (lch Oklch) ToRGBA() RGBA {
	L := lch.Lightness
	a := lch.Chroma * math.Cos(lch.Hue)
	b := lch.Chroma * math.Sin(lch.Hue)

	l := L + 0.3963377774*a + 0.2158037573*b
	m := L - 0.1055613458*a - 0.0638541728*b
	s := L - 0.0894841775*a - 1.2914855480*b

	l = l * l * l
	m = m * m * m
	s = s * s * s

	linearRed := 4.0767416621*l - 3.3077115913*m + 0.2309699292*s
	linearGreen := -1.2684380046*l + 2.6097574011*m - 0.3413193965*s
	linearBlue := -0.0041960863*l - 0.7034186147*m + 1.7076147010*s

	// Convert from linear sRGB to sRGB
	// https://bottosson.github.io/posts/colorwrong/#what-can-we-do
	if linearRed >= 0.0031308 {
		linearRed = 1.055*math.Pow(linearRed, 1.0/2.4) - 0.055
	} else {
		linearRed = 12.92 * linearRed
	}
	if linearGreen >= 0.0031308 {
		linearGreen = 1.055*math.Pow(linearGreen, 1.0/2.4) - 0.055
	} else {
		linearGreen = 12.92 * linearGreen
	}
	if linearBlue >= 0.0031308 {
		linearBlue = 1.055*math.Pow(linearBlue, 1.0/2.4) - 0.055
	} else {
		linearBlue = 12.92 * linearBlue
	}

	red := uint8(math.Round(linearRed * 255))
	green := uint8(math.Round(linearGreen * 255))
	blue := uint8(math.Round(linearBlue * 255))
	alpha := uint8(math.Round(lch.Alpha * 255))

	return RGBA{red, green, blue, alpha}
}

func (lch Oklch) ToHex() Hex {
	return lch.ToRGBA().ToHex()
}

func (lch Oklch) String() string {
	alpha := ""
	if lch.Alpha != 1 {
		alpha = fmt.Sprintf(" / %.3f%%", lch.Alpha*100)
	}
	hue := fmt.Sprintf("%f", lch.Hue)
	if lch.Hue == 0 {
		hue = "none"
	}
	return fmt.Sprintf("oklch(%v%% %v %s%s)", lch.Lightness*100, lch.Chroma, hue, alpha)
}
