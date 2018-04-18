// +build ignore

package main

import (
	"fmt"
	"time"

	ansi "github.com/k0kubun/go-ansi"
)

func main() {
	fmt.Print("1111\n")
	fmt.Print("2222\n")
	fmt.Print("3333\n")
	fmt.Print("4444")

	ansi.EraseInLine(1)

	time.Sleep(1 * time.Second)
	fmt.Print("0")

}
