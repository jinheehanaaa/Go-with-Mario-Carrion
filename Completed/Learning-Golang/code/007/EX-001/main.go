// https://github.com/MarioCarrion/videos/blob/8b70c807d63b854cde4dcbf1a1111b856a449cd5/2021/08/12/introductions-concurrency-patterns/01-goroutines/main.go

package main

import "fmt"

// No result b/c..?
// We need "wait" for goroutine
func main() {
	go hello()
}

func hello() {
	fmt.Println("it's most likely you will never see this")
}