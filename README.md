# color

This is a tool to quickly convert between color formats. Currently, it supports:

- OKLCH
- RGB
- Hex

## Build instructions

You can build it from source using Go:

```sh
go install codeberg.org/noclaps/color@latest
```

## Usage

You can use the tool simply by running:

```sh
color <input color> <output format>
```

For example:

```sh
color '#c0ffee' oklch
```

The input color formats are the same as the output formats listed above.
