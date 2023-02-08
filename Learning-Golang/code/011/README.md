
# What is Background Jobs?
- A process in charge of doing some work "behind the scenes"
- Initialized by another "parent" process
- A goroutine launching N goroutines

# EX-001: Background Job
- There is no way to communicate to goroutine (doWork)
- Goroutines can't exit cleanly (But still has USE-CASE)
- USE-CASE: If something fails, everything fails
- - EX: Multiple files to do background calculation & send to another channel using pipeline pattern

# EX-002: Background
- Goroutines process msg through msgC Channel
- Since goroutine is expecting message through channel, we can close the msgC
- and that way, goroutines can exit
