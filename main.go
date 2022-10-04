package main

import (
	"errors"
	"fmt"
	"time"
)

var errRequestFailed = errors.New("Request Failed")

func main() {
	go Count("jootang2 said")
	go Count("Mingji said")
	time.Sleep(time.Second * 5)
}

// func hitURL(url string) error {
// 	fmt.Println("Checking", url)
// 	resp, err := http.Get(url)
// 	if err != nil || resp.StatusCode >= 400 {
// 		fmt.Println(err, resp.StatusCode)
// 		return errRequestFailed
// 	}
// 	return nil
// }

func Count(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "this is golang", i)
		time.Sleep(time.Second)
	}
}
