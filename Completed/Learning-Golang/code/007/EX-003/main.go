// https://github.com/MarioCarrion/videos/blob/8b70c807d63b854cde4dcbf1a1111b856a449cd5/2021/08/12/introductions-concurrency-patterns/03-channels/main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		fmt.Println(time.Now(), "taking a nap")

		time.Sleep(2 * time.Second)

		ch <- "hello"
	}()

	fmt.Println(time.Now(), "waiting for message")

	v := <-ch

	fmt.Println(time.Now(), "received", v)

	// go run main.go
	// waiting for message
	// taking a nap
	// 2 seconds
	// received hello
}