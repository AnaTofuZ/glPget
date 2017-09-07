package main

import (
	"fmt"
	"os"

	"github.com/AnaTofuZ/glPget"
)

func main() {
	cli := glPget.New()

	if err := cli.Run(); err != nil {
		if cli.Trace {
			fmt.Fprintf(os.Stderr, "Error:\n%+v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "Error:\n %v\n", err)
		}
		os.Exit(1)
	}
	os.Exit(0)
}
