package main

import (
	"fmt"

	_ "github.com/MarioCarrion/videos/2021/12/21/28/init/first"
	_ "github.com/MarioCarrion/videos/2021/12/21/28/init/second"
)

// Don't use Init()
// Init is called by File Structure (Alphabetical Order)
// first
// - first.go
// - z.go
func main() {
	fmt.Println("hello from main!")

	// Result:
	// first 1
	// first 2
	// first 3
	// first from 'a'
	// second
	// hello from main!
}
