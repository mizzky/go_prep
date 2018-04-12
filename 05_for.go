// +build ignore

package main

import (
	"fmt"
)

var (
	sum1 = 0
	sum2 = 0
)

func main() {

	// 0から100未満の偶数を累計する part1
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			sum1 += i
		}
	}

	// 0から100未満の偶数を累計する part2
	for i, r := range  {
		fmt.Println(i)
	}

	fmt.Println(sum1)

}
