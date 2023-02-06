# What is Concurrency?
- Concurrency is the composition of independently executing processes
- Concurrency is about dealing with lots of things at once

# What is Parallelism?
- Parallelism is the simultaneous execution of (possibly related) computations
- Parallelism is about doing lots of things at once

# What is a Goroutine?
- Independently executed function

# EX-001: Goroutine requires wait indication


# What is a Channel?
- Mechanism through which we can send and receive values
- <code>ch := make(chan T)</code>
- <code><-</code> indicates the data flow direction
- <code>ch <-</code> v is going into the channel
- <code>v = <- ch</code> channel is sending to  v
- By default sends and receives block until the other side is ready

# Buffered Channel
- Length
- <code>ch := make(chan T, n)</code>
- "Sending" blocks when buffer is full
- "Receiving" blocks when buffer is full

## EX-004
- Channel is sending values with sized buffer
- But there will be a time where your buffer can be full
- So we need to take care of race condition (memory leak)

# Channel & What else
- close => close the channel
- range => pull the value automatically until channel is closed
- select => coordinate with different channel (similar to switch)

## EX-005: Using 2 Channels (select)
- Multiple channels

## EX-006: Using 1 Channels (range)
- Resulat is the same as EX-005
- We will learn when should you use select or range later in this course


# Resources
- [YouTube](https://www.youtube.com/watch?v=-xEycsoGoA8&list=PL7yAAGMOat_F7bOImcjx4ZnCtfyNEqzCy&index=8)
- [Example Code](https://github.com/MarioCarrion/videos/tree/8b70c807d63b854cde4dcbf1a1111b856a449cd5/2021/08/12/introductions-concurrency-patterns)