# Objective

# More
- [README.md](code/012/README.md)

# What is in errgroup?
- Synchronization
- Error propagation
- Context cancelation

## EX-001-a: sync package
- Use <code>waitGroups</code> for goroutines to work together
- Define length of how many goroutines are supposed to be completed for whole process
## EX-001-b: errgroup package
- Use errGroup
- Go Methods: Specify function that runs as goroutine (also returns error if something fails)
- Wait Methods inside func: Propagates errors using context package

## EX-002: errgroup with context package
- We can pass ctx info to goroutines (ex: close cancel goroutine)

# Resources
- [YouTube](https://www.youtube.com/watch?v=KGOgEr7tFKE&list=PL7yAAGMOat_F7bOImcjx4ZnCtfyNEqzCy&index=10)
- [Code Example](https://github.com/MarioCarrion/videos/tree/a2d4d2d8efb306f44c23e157d0ad10636230f3de/2021/09/02)
- [errgroup package](https://pkg.go.dev/golang.org/x/sync/errgroup)
