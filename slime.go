// +build ignore
// ----------------------------------
// this is first cli program..
// date 		: 2018/04/11
// author 	:	mizzky
// ----------------------------------

package main

import (
	"flag"
	"fmt"
)

// define option
var (
	show_version = flag.Bool("version", false, "show version")
)

var (
	version = "v1.10"
)

func main() {
	flag.Parse()

	if *show_version {
		fmt.Printf("version: %s\n", version)
		return
	}
}
