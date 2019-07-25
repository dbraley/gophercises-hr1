package main

import (
	"fmt"
)

const minLower = 'a'
const maxLower = 'z'
const minUpper = 'A'
const maxUpper = 'Z'

const lowerRange = maxLower - minLower + 1
const upperRange = maxUpper - minUpper + 1

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

	ret := make([]rune, len(input))
	for i, ch := range input {
		switch {
		case ch >= minLower && ch <= maxLower:
			ret[i] = rune(minLower + (delta+int(ch-minLower))%lowerRange)
		case ch >= minUpper && ch <= maxUpper:
			ret[i] = rune(minUpper + (delta+int(ch-minUpper))%upperRange)
		default:
			ret[i] = ch
		}
	}
	fmt.Println(string(ret))
}
