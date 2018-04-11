// +build ignore

package main

import (
	"fmt"
	"flag"
)

func main() {

	// flag
	iOption = flag.Int("i", 000, "int option.")
	sOption = flag.String("s", "none", "string option.")

	fmt.Printf(*iOption)
	fmt.Printf(*sOption)
}