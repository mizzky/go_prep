// +build ignore

package main

import (
	"fmt",
	
)

var (

	// 	str_01 =
	// " 	�@�@. �_�F�^ \n
	// 	�@�E���������E  \n
	// 	. ?,,?. ��.�_ \n
	// 	(*�E�ցE)/ \n
	// 	.(�@ � \n
	// 	.���[�i \n
	// "

	// 		str_02 =
	// 	"   	�@�@. �_�F�^ \n
	// 			�@�E���������E  \n
	// 			. ?,,?. ��.�_ \n
	// 			(*�E�ցE)/ \n
	// 			.(�@ � \n
	// 			.���[�i \n
	// 	"

	str01 = [6]string{
		"     . �_�F�^ \r",
		"    �E���������E  \r",
		"  ?,,?. ��.�_ \r",
		" (*�E�ցE)/ \r",
		".(�@ � \r ",
		".���[�i \r"}

	str02 = [6]string{
		"       . �_�F�^ \r",
		"      �E���������E  \r",
		"    ?,,?. ��.�_ \r",
		"   (*�E�ցE)/ \r",
		"  .(�@ � \r ",
		"  .���[�i \r"}
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
