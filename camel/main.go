package main

import "fmt"

func main() {
	var input string
	if _, err := fmt.Scanf("%s\n", &input); err != nil {
		panic(err)
	}
	answer := 1

	min, max := 'A', 'Z'
	for _, ch := range input {
		if ch >= min && ch <= max {
			answer++
		}
	}

	fmt.Println(answer)
}
