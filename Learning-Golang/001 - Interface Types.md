# Types
## Basic Types
- strings
- numeric
- bool

## Composite Types
- pointer
- function
- map
- slice
- array
- channel
- interface

# Functions and Methods
## EX1
```go
package main

import "fmt"

type Priority int8

func (p Priority) String() string {
    switch int8(p) {
        case 0:
            return "low"
        case 1:
        return "high"
    }

    return "unknown"
}

func main() {
    var p Priority =1
    fmt.Println(p.String()) // Method
    fmt.Println(Priority.String(p)) // Function
    fmt.Println(p)
}
```

# Interface Type
- A set of method signatures
- Can hold any value that implements those methods
- EX1 works because of Stringer interface














