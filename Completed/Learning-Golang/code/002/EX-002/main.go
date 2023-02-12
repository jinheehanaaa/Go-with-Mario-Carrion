package main

import "fmt"

type Interface interface {
	Method()
}

type String struct {
	Value string
}

func (s *String) Method() {
	fmt.Printf("s = %v\n", s)
}

type Integer int

func (i Integer) Method() {}

// Interface value is nil if both value & concrete type are nil
func main() {
	var str *String
	fmt.Printf("str= %v\n", str) // 1
	str.Method()                 // 2

	(*String).Method(str) // 3 (Syntactic sugar)

	var iface Interface = str
	iface.Method() // 4
}

// go run main.go
// str= <nil>
// s= <nil>
// s= <nil>
// s= <nil>
