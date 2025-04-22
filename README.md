# color

This is a CLI tool to quickly convert between color formats. It's based on the [Color.js](https://colorjs.io) library and supports practically every format supported by CSS.

## Installation

You can download it from [Releases](https://github.com/noClaps/color/releases), or from Homebrew:

```sh
brew install noclaps/tap/color
```

or you can build it from source:

```sh
git clone https://github.com/noClaps/color.git && cd color
bun install
bun run build
```

## Usage

```
USAGE: color <color> <format> [--list-formats]

ARGUMENTS:
  <color>               The color that you would like to convert.
  <format>              The format that you would like to convert to.

OPTIONS:
  --list-formats, -f    List all the available formats and exit.
  --help, -h            Display this help message and exit.
```

You can use the tool simply by running:

```sh
color "#ff732e" oklch # oklch(71.475% 0.18776 43.447)
```

The supported color formats can be listed using the `--list-formats` or `-f` flag:

```sh
color --list-formats
color -f
```

You can view the help by using `--help` or `-h`:

```sh
color --help
color -h
```
