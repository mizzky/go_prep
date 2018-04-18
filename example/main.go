package main

import (
	"fmt"

	prep "github.com/mizzky/cli_prep/sum"
)

func main() {
	ans := prep.Sum(1, 3)
	fmt.Print(ans)
}
