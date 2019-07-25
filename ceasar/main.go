package main

import (
	"fmt"
	"strings"
)

const lower = "abcdefghijklmnopqrstuvwxyz"
const upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	var length, delta int
	var input string

	if _, err := fmt.Scanf("%d\n", &length); err != nil {
		panic(err)
	}
	if _, err := fmt.Scanf("%s\n", &input); err != nil {
		panic(err)
	}
	if _, err := fmt.Scanf("%d\n", &delta); err != nil {
		panic(err)
	}

	ret := ""
	for _, ch := range input {
		switch {
		case strings.ContainsRune(lower, ch):
			ret += string(rotate(ch, delta, []rune(lower)))
		case strings.ContainsRune(upper, ch):
			ret += string(rotate(ch, delta, []rune(upper)))
		default:
			ret += string(ch)
		}
	}
	fmt.Println(ret)
}

func rotate(s rune, delta int, key []rune) rune {
	idx := strings.IndexRune(string(key), s)
	if idx < 0 {
		panic("idx < 0")
	}
	idx = (idx + delta) % len(key)
	return key[idx]
}
