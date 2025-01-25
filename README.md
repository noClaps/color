# color

This is a CLI tool to quickly convert between color formats. It's based on the [Color.js](https://colorjs.io) library and supports practically every format supported by CSS.

## Build instructions

You'll need [Bun](https://bun.sh) to build this project.

1. Clone the repository.

   ```sh
   git clone https://gitlab.com/noClaps/color.git
   cd color
   ```

2. Build the project.

   ```sh
   bun install
   make build
   ```

You can then run it using `./color`

## Usage

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
