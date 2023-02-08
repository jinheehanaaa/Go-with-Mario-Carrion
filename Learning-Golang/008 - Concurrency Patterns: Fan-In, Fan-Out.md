# Objective

# More
- [README.md](code/012/README.md)

# What is the Fan-In pattern?
- Consolidation of multiple channels into one channel by multiplexing each received value



## EX-001: Fan-In Pattern (Merging 2 files)
### merge1: sync package
- Use <code>WaitGroup</code> to identify how many times I should be receiving the value back
- Result depens on the speed of goroutine & your system (Random)
### merge2: buffer channel
- Use separate buffer channel to identify when to stop the goroutine
- No sync package

# What is the Fan-Out pattern?
- Breakup of one channel into multiple ones by distributing each value.

## EX-002: Fan-Out Pattern (1 Channel into multiple one)



# Resources
- [YouTube](https://www.youtube.com/watch?v=x6vBvgKGvxU&list=PL7yAAGMOat_F7bOImcjx4ZnCtfyNEqzCy&index=9)
- [Code Example](https://github.com/MarioCarrion/videos/tree/3107ff408e0db59b5e9ae07412460375aeb8786a/2021/08/19)