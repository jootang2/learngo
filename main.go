package main

import (
	"errors"
	"fmt"
	"time"
)

var errRequestFailed = errors.New("Request Failed")

func main() {
	channel := make(chan bool)
	people := [2]string{"juhwan", "mingji"}
	for _, person := range people {
		go isHappy(person, channel)
	}
	result := <- channel
	fmt.Println(result)

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

func isHappy(person string, channel chan bool) {
	time.Sleep(time.Second * 5)
	channel <- true
}
