# 1. context should be the first parameter

# 2. Return early

# 3. Avoid init(). (99.9% of the time)
- init() allows you to initialize global variables or functions
- But, anything that has global state is hard to test
- gives you errors when dealing with concurrency

# 4. Comments are for humans and machines
- Following comment will not be included in documentation in Godev website
```go
// Provider ...
// nolint: revive
```

# 5. Avoid unnecessary functions
- Refer to Functional Option

# Resources
- [YouTube](https://www.youtube.com/watch?v=lfQ4qLcE3Bo&list=PL7yAAGMOat_F7bOImcjx4ZnCtfyNEqzCy&index=20)
- [Code Example](https://github.com/MarioCarrion/videos/tree/301200a7bc2c85cc4f66880892234d645b377279/2021/12/28)