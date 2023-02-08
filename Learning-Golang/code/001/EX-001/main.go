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
	var p Priority = 1
	fmt.Println(p.String())         // Method
	fmt.Println(Priority.String(p)) // Function
	fmt.Println(p)
}
