package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	x := "texts"
	data := "♥"
	fmt.Println(len(data))
	fmt.Println(utf8.ValidString(x))
	xbyte := []byte(x)
	xbyte[0] = 'T'

	fmt.Println(xbyte)

	m := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	for k, v := range m {
		fmt.Println(k, v)
	}
	var b uint8 = 0x02
	fmt.Printf("%08b (NOT B)\n", ^b)
}
