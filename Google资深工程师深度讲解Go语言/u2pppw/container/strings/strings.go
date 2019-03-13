package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "好久不见"
	fmt.Println(len(s))
	fmt.Println([]byte(s))

	// int32 == rune
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()


	fmt.Println(utf8.RuneCountInString(s))


	bytes := []byte(s)

	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()


	for i, ch := range []rune(s) {
		fmt.Printf("%d ：%c\n", i, ch)
	}
	fmt.Println()


	// strings.
}
