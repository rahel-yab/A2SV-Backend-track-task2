package main

import (
	"fmt"
	"strings"
	"unicode"
)

func palindrome(word string) bool{
	left := 0
	right := len(word) - 1

	for left <= right{
		if word[left] != word[right]{
			return false
		}
		left ++
		right --	
	}
	return true

}
func cleanInput(word string) string {
    var cleaned strings.Builder
    for _, r := range word {
        if unicode.IsLetter(r) || unicode.IsDigit(r) {
            cleaned.WriteRune(unicode.ToLower(r))
        }
    }
    return cleaned.String()
}


func main(){
	val := "A man, a plan, a canal: Panama"
	cleaned := cleanInput(val)
	fmt.Println(palindrome(cleaned))
}