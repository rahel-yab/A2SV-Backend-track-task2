package main

import (
	"reflect"
	"testing"
)

func TestFrequency(t *testing.T) {
    input := "go go gone"
    expected := map[string]int{"go": 2, "gone": 1}

    result := frequency(input)
    if !reflect.DeepEqual(result, expected) {
        t.Errorf("Expected %v, got %v", expected, result)
    }
}

