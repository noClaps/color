package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"codeberg.org/noclaps/color/internal/color"
)

func main() {
	flag.Parse()

	colorArg := flag.Arg(0)
	format := flag.Arg(1)

	format = strings.ToLower(format)

	if format != "oklch" && format != "rgb" && format != "hex" {
		log.Fatalf("Unsupported format: `%s`. Supported formats are: 'oklch', 'rgb', 'hex'", format)
		return
	}

	if len(colorArg) > 3 && len(colorArg) < 10 && colorArg[0] == '#' {
		hex, err := color.NewHex(colorArg)
		if err != nil {
			log.Fatalln(err)
		}

		switch format {
		case "oklch":
			fmt.Println(hex.ToOklch())
		case "rgb":
			fmt.Println(hex.ToRGBA())
		case "hex":
			fmt.Println(hex)
		}
		return
	}

	if len(colorArg) > 7 && colorArg[0:6] == "oklch(" && colorArg[len(colorArg)-1] == ')' {
		oklch, err := color.NewOklch(colorArg)
		if err != nil {
			log.Fatalln(err)
		}

		switch format {
		case "oklch":
			fmt.Println(oklch)
		case "rgb":
			fmt.Println(oklch.ToRGBA())
		case "hex":
			fmt.Println(oklch.ToHex())
		}
		return
	}

	if len(colorArg) > 5 && colorArg[0:4] == "rgb(" && colorArg[len(colorArg)-1] == ')' {
		rgb, err := color.NewRGBA(colorArg)
		if err != nil {
			log.Fatalln(err)
		}

		switch format {
		case "oklch":
			fmt.Println(rgb.ToOklch())
		case "rgb":
			fmt.Println(rgb)
		case "hex":
			fmt.Println(rgb.ToHex())
		}
		return
	}

	log.Fatalf("Invalid input: `%v`", colorArg)
}
