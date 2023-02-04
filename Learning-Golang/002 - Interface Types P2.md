# Interface value
- A tuple of a value and a concrete type
- An Interface value is nil if:
- - both its value and concrete type are nil


# EX 1: Intro to interface

```go
package main

import (
    'fmt'
)

type Interface interface {
    Method()
}

type String struct {
    Value string
}

func (s *String) Method() {}

type Integer int

func (i Integer) Method() {}

func main() {
    var iface Interface

    iface = &String{"hello world"}
    fmt.Printf("Value: %v, Type: %T\n", iface, iface)

    iface = Integer(100)
    fmt.Printf("Value: %v, Type: %T\n", iface, iface)
}

// go run main.go
// Value: &{hello world}, Type: *main.String
// Value: 100, Type: main.Integer
```

# EX 2: Calling interface values with nil?

```go
package main

import (
    'fmt'
)

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

func main() {
    var str *String
    fmt.Printf("str= %v\n", str)
    str.Method()

    (*String).Method(str) // Syntactic sugar

    var iface Interface = str
    iface.Method
}

// go run main.go
// str= <nil>
// s= <nil>
// s= <nil>
// s= <nil> 
```

# Interface{}
- Specifies zero methods
- May hold any values of any type
- Used by code that handles unknown types

## Examples of interface{}
- Functions in fmt
- Functions in encoding/json
- Types in database/sql

## Using interface{}
- Type Assertions
- Type Switches


## EX 3: Type Assertion
```go
package main

import "fmt"

type Interface interface {
	Method()
}

type Integer int

func (i Integer) Method() {}

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
```


# Resources
- [YouTube](https://www.youtube.com/watch?v=m14ob5dCLag&list=PL7yAAGMOat_F7bOImcjx4ZnCtfyNEqzCy&index=3)
- [Code Examples](https://github.com/MarioCarrion/videos/blob/main/2021/07/interfaces-part-2/03-type-assertion-switches/main.go)