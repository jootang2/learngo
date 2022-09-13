package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (lenght int, uppercase string) {
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)

}

// test
func main() {
	totalLenght, up := lenAndUpper("joohwan")
	fmt.Println(totalLenght, up)
}
