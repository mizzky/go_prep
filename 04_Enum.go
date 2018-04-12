// +build ignore

package main

import "fmt"

const (
	朝 = "おはよう"
	昼 = "こんにちは"
	夕 = "こんばんは"
)

func 挨拶(m string) {
	fmt.Println(m)
}

func main() {
	挨拶(昼)
}
