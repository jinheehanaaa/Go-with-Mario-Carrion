package main

import (
	"dont-return-interfaces/producer1"
	"dont-return-interfaces/producer2"
	"fmt"
)

// This doesn't force your users to use concrete type
type Stringer interface {
	String() string
}

// Producer1 returns interface

// Producer2 returns struct

func main() {
	p1 := producer1.New()
	fmt.Println("name", p1.Produce())

	p2 := producer2.New()
	fmt.Println("car", p2.Produce())
}
