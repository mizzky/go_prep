package main

import (
	"fmt"
	"os"
)


func main() {

	// require 3 element args
	if len(os.Args) != 3 {
		fmt.Printf(" invalid args.")
		os.Exit(1)
	}

	// param1
	fmt.Printf("arg1: %s\n", os.Args[1])
	fmt.Printf("arg2: %s\n", os.Args[2])
	
}