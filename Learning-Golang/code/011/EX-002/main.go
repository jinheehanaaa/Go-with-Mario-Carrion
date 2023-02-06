package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// # of Workers (5) => Process
// Buffer Channel Size (10) => How many messages can this channel queue?
// go build main.go
// kill -s TERM 415479
// for i in {1..15}; do kill -s TERM 415479; done
func main() {
	fmt.Println("Process ID", os.Getpid())

	s := NewScheduler(5, 10)

	s.ListenForWork()

	fmt.Println("ready")

	<-waitToExit()

	s.Exit()

	fmt.Println("exiting")
}

// Listen for Interrupt signal
func waitToExit() <-chan struct{} {
	runC := make(chan struct{}, 1)

	sc := make(chan os.Signal, 1)

	signal.Notify(sc, os.Interrupt)

	go func() {
		defer close(runC)

		<-sc
	}()

	return runC
}

//-

type Scheduler struct {
	workers   int
	msgC      chan struct{}  // Channel for receving the event
	signalC   chan os.Signal // Channel for stopping process
	waitGroup sync.WaitGroup
}

func NewScheduler(workers, buffer int) *Scheduler {
	return &Scheduler{
		workers: workers,
		msgC:    make(chan struct{}, buffer),
		signalC: make(chan os.Signal, 1),
	}
}

func (s *Scheduler) ListenForWork() {
	go func() { // 1) Listen for messages to process

		signal.Notify(s.signalC, syscall.SIGTERM)

		for {
			<-s.signalC

			s.msgC <- struct{}{} // 2) Send to processing channel
		}
	}()

	s.waitGroup.Add(s.workers)

	for i := 0; i < s.workers; i++ {
		i := i
		go func() {
			for {
				select {
				case _, open := <-s.msgC: // 3) Wait for messages to process
					if !open { // closed, exiting
						fmt.Printf("%d closing\n", i+1)
						s.waitGroup.Done()

						return
					}

					fmt.Printf("%d<- Processing\n", i)
				}
			}
		}()
	}
}

func (s *Scheduler) Exit() {
	close(s.msgC)
	s.waitGroup.Wait()
}
