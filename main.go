package main

import "fmt"

// test
func main() {
	a := 2         //a 주소에 담긴 값 : 2
	b := &a        //b : a 주소 값
	*b = 4         //*b : b 값 (a 주소 값)을 4로 변경 => a 주소에 담긴 값 4로 변경
	fmt.Println(a) // => 출력 : 4
}
