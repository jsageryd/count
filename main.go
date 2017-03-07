package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	n, err := io.Copy(os.Stdout, os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "count: %v\n", err)
	}
	fmt.Fprintf(os.Stderr, "count: %d bytes\n", n)
}
