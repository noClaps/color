package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/noclaps/applause"
	"github.com/noclaps/color/internal/color"
	"github.com/noclaps/color/internal/logger"
)

type Args struct {
	Color  string `help:"The color that you would like to convert."`
	Format string `help:"The format that you would like to convert to. Supported formats are: 'oklch', 'rgb', 'hex'."`
}

func main() {
	args := Args{}
	if err := applause.Parse(&args); err != nil {
		logger.Errorf("%v", err)
		logger.Logf("%s", applause.Usage)
		os.Exit(1)
	}

	args.Format = strings.ToLower(args.Format)

	if args.Format != "oklch" && args.Format != "rgb" && args.Format != "hex" {
		logger.Fatalf("Unsupported format: `%s`. Supported formats are: 'oklch', 'rgb', 'hex'", args.Format)
		return
	}

	if len(args.Color) > 3 && len(args.Color) < 10 && args.Color[0] == '#' {
		hex, err := color.NewHex(args.Color)
		if err != nil {
			logger.Fatalf("%v", err)
		}
		switch args.Format {
		case "oklch":
			fmt.Println(hex.ToOklch())
		case "rgb":
			fmt.Println(hex.ToRGBA())
		case "hex":
			fmt.Println(hex)
		default:
			logger.Fatalf("Invalid format: `%v`", args.Format)
		}

		return
	}

	if len(args.Color) > 7 && args.Color[0:6] == "oklch(" && args.Color[len(args.Color)-1] == ')' {
		oklch, err := color.NewOklch(args.Color)
		if err != nil {
			logger.Fatalf("%v", err)
		}

		switch args.Format {
		case "oklch":
			fmt.Println(oklch)
		case "rgb":
			fmt.Println(oklch.ToRGBA())
		case "hex":
			fmt.Println(oklch.ToHex())
		default:
			logger.Fatalf("Invalid format: `%v`", args.Format)
		}
		return
	}

	if len(args.Color) > 5 && args.Color[0:4] == "rgb(" && args.Color[len(args.Color)-1] == ')' {
		rgb, err := color.NewRGBA(args.Color)
		if err != nil {
			logger.Fatalf("%v", err)
		}

		switch args.Format {
		case "oklch":
			fmt.Println(rgb.ToOklch())
		case "rgb":
			fmt.Println(rgb)
		case "hex":
			fmt.Println(rgb.ToHex())
		default:
			logger.Fatalf("Invalid format: `%v`", args.Format)
		}
		return
	}

	logger.Fatalf("Invalid input: `%v`", args.Color)
}
