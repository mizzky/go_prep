// +build ignore

package main

import (
	"fmt",
	
)

var (

	// 	str_01 =
	// " 	@@. _F^ \n
	// 	@EĽĽĽĽĽĽE  \n
	// 	. ?,,?. ż._ \n
	// 	(*EÖE)/ \n
	// 	.(Â@ É \n
	// 	.ľ[i \n
	// "

	// 		str_02 =
	// 	"   	@@. _F^ \n
	// 			@EĽĽĽĽĽĽE  \n
	// 			. ?,,?. ż._ \n
	// 			(*EÖE)/ \n
	// 			.(Â@ É \n
	// 			.ľ[i \n
	// 	"

	str01 = [6]string{
		"     . _F^ \r",
		"    EĽĽĽĽĽĽE  \r",
		"  ?,,?. ż._ \r",
		" (*EÖE)/ \r",
		".(Â@ É \r ",
		".ľ[i \r"}

	str02 = [6]string{
		"       . _F^ \r",
		"      EĽĽĽĽĽĽE  \r",
		"    ?,,?. ż._ \r",
		"   (*EÖE)/ \r",
		"  .(Â@ É \r ",
		"  .ľ[i \r"}
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
