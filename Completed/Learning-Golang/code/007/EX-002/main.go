package main

// https://github.com/MarioCarrion/videos/blob/8b70c807d63b854cde4dcbf1a1111b856a449cd5/2021/08/12/introductions-concurrency-patterns/02-channels/main.go
import "fmt"

// XXX: This program panics because there is no goroutine
// outside of `main` interacting with the `ch` channel:
//
// fatal error: all goroutines are asleep - deadlock!

func main() {
	var ch chan int
	ch = make(chan int)
	// Two lines above can be also replaced with
	// ch := make(chan int)

	ch <- 10

	v := <-ch
	fmt.Println("received", v)
}