package main

import (
	"reflect"
	"testing"
)

func TestPalindrome(t *testing.T) {
    input := "hello world"
    expected := false
	result := palindrome(input)
	
    if !reflect.DeepEqual(result, expected) {
        t.Errorf("Expected %v, got %v", expected, result)
    }
}

