# What is Background Jobs?
- A process in charge of doing some work "behind the scenes"
- Initialized by another "parent" process
- A goroutine launching N goroutines

# EX-001: Background Job
- There is no way to communicate to goroutine (doWork)
- Goroutines can't exit cleanly
- USE-CASE: If something fails, everything fails
- - EX: Multiple files to do background calculation & send to another channel using pipeline pattern

# EX-002: Background
- Goroutines process msg through msgC
- Since goroutine is expecting message through channel, we can close the msgC
- and that way, goroutines can exit

# Resources
- [YouTube](https://www.youtube.com/watch?v=sKvFXAkQqXY&list=PL7yAAGMOat_F7bOImcjx4ZnCtfyNEqzCy&index=12)
- [Code Example](https://github.com/MarioCarrion/videos/tree/7a00f3d7fc1573985d4403d122a230a2554acb93/2021/10/01)