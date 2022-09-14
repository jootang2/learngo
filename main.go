package main

import (
	"fmt"
)

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge <= 18 {
		return false
	}
	return true
}

// test
func main() {
	fmt.Print(canIDrink(16))
}
