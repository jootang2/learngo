package main

import "fmt"

// test
func main() {
	jootang2 := map[string]string{"name": "jootang2" , "age" : "27"}
	for _,value := range jootang2 {
		fmt.Println(value)
	}
}
