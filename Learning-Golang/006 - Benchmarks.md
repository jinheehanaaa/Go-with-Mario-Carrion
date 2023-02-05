# Which Benchmark is fastest for reading string?
- Sprintf (Exclude this becuase it's too slow)
- Bytes
- StringsBuilder

# Benchmark
## Basic
- <code>go test -bench=.</code> for running all bench functions
- <code>go test -bench=. -benchtime=1000x<</code> for x amount of time
- <code>go test -bench=. -benchtime=1s<</code> for x amount of second
- <code>go test -bench=. -benchtime=2s -benchmem<</code> with memory
- - Sometimes you need to make decision between time spent vs. memory usage

## 1 Run
- <code>go test -run=BenchmarkConcat -bench=. -benchtime=2s -benchmem</code>



# Packages
- benchstat <code>go install golang.org/x/perf/cmd/benchstat@latest</code>



# Resources
- [YouTube](https://www.youtube.com/watch?v=u6dpEuJ7tB8&list=PL7yAAGMOat_F7bOImcjx4ZnCtfyNEqzCy&index=7)
