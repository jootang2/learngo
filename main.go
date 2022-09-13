package main

import (
	"fmt"
)

func multiply(a, b int) int {
	return a * b
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func repeatMe(words ...string) {
	fmt.Println(words)

}

// test
func main() {
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
