# 1. Prefer table driven tests
- Testing is much more easy if you add/modify actual code

# 2. Avoid Assert Tests
- Just test directly where you are supposed to be testing instead of asserting
## Case 1: Assert
- github.com/stretchr/testify
```go
assert.Equal(t, tt.output. got)
```
## Case 2: Test with Fatalf
- No package
```go
if got != tt.output {
		t.Fatalf("got %v, wanted %v", got, tt.output)
}
```
# 3. Prefer not using interface{} (99.9% of the time)

# 4. Don't return interfaces

# 5. Avoid stuttering


# Resources
- [YouTube](https://www.youtube.com/watch?v=qCg-FIkcJZw&list=PL7yAAGMOat_F7bOImcjx4ZnCtfyNEqzCy&index=21)
- [Example Code](https://github.com/MarioCarrion/videos/tree/5921bdac1dcaf92bae65807df447d3dbd833fe65/2022/01/21)
