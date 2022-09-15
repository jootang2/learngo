package main

import (
	"fmt"
)

func canIDrink(age int) bool {
	switch koreanAge := age + 2; koreanAge{
	case 19:
		return false
	case 20:
		return true
	}
	return false
}

// test
func main() {
	fmt.Print(canIDrink(19))
}
