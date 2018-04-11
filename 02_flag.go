// +build ignore

package main

import (
	"flag"
	"fmt"
)

func main() {

	// flag
	iOption := flag.Int("i", 000, "int option.")
	sOption := flag.String("s", "none", "string option.")

	flag.Parse()
	fmt.Println(*iOption)
	fmt.Println(*sOption)
}
