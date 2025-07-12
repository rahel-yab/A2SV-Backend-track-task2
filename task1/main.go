package main

import (
	"fmt"
	"strings"
	"unicode"
)

func cleanInput(input string) string {
    var curr strings.Builder
    for _, r := range input {
        if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) {
            curr.WriteRune(unicode.ToLower(r))
        }
    }
    return curr.String()
}

func frequency(input string) map[string]int {
    dict := make(map[string]int)
    cleaned := cleanInput(input)
    words := strings.Fields(cleaned)
    for _, word := range words {
        dict[word]++
    }
    return dict
}

func main() {
    fmt.Println(frequency("I am learning Go. go GO!"))
}