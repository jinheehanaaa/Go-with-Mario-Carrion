# Benchmark
## Basic Bench Command
- <code>go test -bench=.</code> for running all bench functions
- <code>go test -bench=. -benchtime=1000x<</code> for x amount of time
- <code>go test -bench=. -benchtime=1s<</code> for x amount of second
- <code>go test -bench=. -benchtime=2s -benchmem<</code> with memory
- - Sometimes you need to make decision between time spent vs. memory usage

## Compare with benchstat
1. Sprintf: <code>go test -count=5 -run=BenchmarkConcat -bench=. -benchmem | tee current.txt</code>
2. StringsBuilder: <code>go test -count=5 -run=BenchmarkConcat -bench=. -benchmem | tee new.txt</code>
3. Compare: <code>benchstat new.txt current.txt</code>
- - You can compare new and current

### Conclusion
- Only test for things that you can have in your system's memory
- Do not use benchmark for any I/O Operation
- EX: DB Connection, Remote Connection

# Packages
- benchstat: <code>go install golang.org/x/perf/cmd/benchstat@latest</code>