# Builds the color binary
build:
	@bun build index.ts --compile --minify --outfile color
	@echo "Built color"

# Installs the color binary to ~/.local/bin
install: build
	@install ./color ~/.local/bin
	@rm ./color
	@echo "Installed color to ~/.local/bin"
