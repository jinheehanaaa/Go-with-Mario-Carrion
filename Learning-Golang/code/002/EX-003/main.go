package main

import "fmt"

type Interface interface {
	Method()
}

type Integer int

func (i Integer) Method() {}

// TYPE ASSERTION
// panic if you don't use ok
func main() {
	// Topic 1
	var iface interface{} = Integer(100)

	t, ok := iface.(Integer)
	fmt.Printf("OK? %t, Value %v, Type %T\n", ok, t, t)

	iface = "hello"

	t, ok = iface.(Integer)
	fmt.Printf("OK? %t, Value %v, Type %T\n", ok, t, t)

	// Topic 2
	describe("hello")
	describe(Integer(100))
	describe(10)
}

// Use-Case: You might want to deal with different kind of errors
func describe(i interface{}) {
	switch v := i.(type) {
	case Integer:
		fmt.Printf("int %d\n", v)
	case string:
		fmt.Printf("string %s\n", v)
	default:
		fmt.Printf("unknown %T - %v\n", v, v)
	}
}
