package main

import (
	"fmt"

	"github.com/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first" : "First Word"}
	definition, err := dictionary.Search("first") 
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
