// +build ignore

package main

import (
	"fmt"
	"time"
)

const (
	str1 = `
　　. ＼：／ 
　・･･･☆･･･・ 
. ⋀,,⋀. ∩.＼ 
(*・ω・)/ 
.(つ　 ﾉ 
.しーＪ
`
	str2 = `
	　　. ＼：／ 
	　・･･･☆･･･・ 
	. ⋀,,⋀. ∩.＼ 
	(*・ω・)/ 
	.(つ　 ﾉ 
	.しーＪ
`
	str3 = `
		　　. ＼：／ 
		　・･･･☆･･･・ 
		. ⋀,,⋀. ∩.＼ 
		(*・ω・)/ 
		.(つ　 ﾉ 
		.しーＪ
`
)

func main() {

	for i := 1; i <= 100; i++ {
		fmt.Printf(".....%d％\r", i)
		time.Sleep(50 * time.Millisecond)
	}

}
