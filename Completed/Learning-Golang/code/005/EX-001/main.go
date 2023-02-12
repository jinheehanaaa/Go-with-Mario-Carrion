// https://github.com/MarioCarrion/videos/blob/563b41660420a5e77d25157f1d4798f343d12d22/2021/05/07/main_deadline.go

package main

import (
	"context"
	"fmt"
	"time"
)

// PLAY WITH THIS VALUE!
// const shortDuration = 300 * time.Millisecond
const shortDuration = 2 * time.Second

// WithTimeout
func main() {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer cancel()

	// We use select for using goroutine
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}
}
