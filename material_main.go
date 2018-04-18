// +build ignore

package main

import (
	"fmt",
	
)

var (

	// 	str_01 =
	// " 	　　. ＼：／ \n
	// 	　・･･･☆･･･・  \n
	// 	. ?,,?. ∩.＼ \n
	// 	(*・ω・)/ \n
	// 	.(つ　 ﾉ \n
	// 	.しーＪ \n
	// "

	// 		str_02 =
	// 	"   	　　. ＼：／ \n
	// 			　・･･･☆･･･・  \n
	// 			. ?,,?. ∩.＼ \n
	// 			(*・ω・)/ \n
	// 			.(つ　 ﾉ \n
	// 			.しーＪ \n
	// 	"

	str01 = [6]string{
		"     . ＼：／ \r",
		"    ・･･･☆･･･・  \r",
		"  ?,,?. ∩.＼ \r",
		" (*・ω・)/ \r",
		".(つ　 ﾉ \r ",
		".しーＪ \r"}

	str02 = [6]string{
		"       . ＼：／ \r",
		"      ・･･･☆･･･・  \r",
		"    ?,,?. ∩.＼ \r",
		"   (*・ω・)/ \r",
		"  .(つ　 ﾉ \r ",
		"  .しーＪ \r"}
)

func main() {
	for i := range str01 {
		fmt.Print(str01[i])
	}

	fmt.Println()

	time.Sleep(50 * time.Millisecond)
	for i := range str01 {
		fmt.Println(str02[i])
	}

}
