package main

import (
	"fmt"
	"strings"
)

func frequency(input string) map[string]int {
	dict := make(map[string]int)
	words := strings.Fields(input)
	for _, word := range words {
		dict[word] ++
	}
	return dict
}

func main() {
	fmt.Println(frequency("I am learning go"))
}