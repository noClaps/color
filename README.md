# color

This is a tool to quickly convert between color formats. Currently, it supports:

- OKLCH
- RGB
- Hex

## Build instructions

You can build it from source using Go:

```sh
go install github.com/noclaps/color@latest
```

## Usage

```
USAGE: color <color> <format>

ARGUMENTS:
  <color>           The color that you would like to convert.
  <format>          The format that you would like to convert to. Supported formats are: 'oklch',
                    'rgb', 'hex'.

OPTIONS:
  -h, --help        Display this help and exit.
```

You can use the tool simply by running:

```sh
color '#c0ffee' oklch
```

The input color formats are the same as the output formats listed above.

You can view the help by using `-h` or `--help`:

```sh
color -h
color --help
```
