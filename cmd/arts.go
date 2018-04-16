package cmd

import "fmt"

// ArtClose is asciiart
func ArtClose() {
	str := `
	┻┳| 閉めますよ |┳┻ 
	┳┻|_∧ 　 　 ∧_|┻┳ 
	┻┳|･ω･) 　 (•ω•|┳┻ 
	┳┻|⊂ﾉ 　 　 \つ|┻┳ 
	┻┳|∪　 　 　 し|┳┻
	`
	fmt.Println(str)
}

// ArtFav is asciiart
func ArtFav() {
	str := `
	　　. ＼：／ 
	　・･･･☆･･･・ 
	. ⋀,,⋀. ∩.＼ 
	(*・ω・)/ 
	.(つ　 ﾉ 
	.しーＪ
	`
	fmt.Println(str)
}
