package main

import "fmt"

type person struct {
	name string
	age int
	favFood []string
}
// test
func main() {
	favFood := []string{"A","B"}
	joohwan := person{name:"joohwan", age:27, favFood:favFood}
	fmt.Println(joohwan.name)
}
