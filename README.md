# color

This is a tool to quickly convert between color formats. Currently, it supports:

- OKLCH
- RGB
- Hex

## Build instructions

You'll need [Go](https://go.dev).

Clone the repository:

```sh
git clone https://github.com/noClaps/color.git
cd color
```

Build and run the tool:

```sh
go build -o color
./color '#c0ffee' oklch
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
