package logger

import (
	"fmt"
	"os"
)

func Errorf(format string, a ...any) {
	format = "\033[31mERROR:\033[0m " + format + "\n"
	fmt.Fprintf(os.Stderr, format, a...)
}

func Fatalf(format string, a ...any) {
	format = "\033[31mERROR:\033[0m " + format + "\n"
	fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(1)
}

func Logf(format string, a ...any) {
	format = format + "\n"
	fmt.Fprintf(os.Stderr, format, a...)
}
