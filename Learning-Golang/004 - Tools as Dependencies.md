# DIRENV
- You can use direnv to setup the environment differently
- - Use your desired version of binary
- - No collision between the versions

# Package Version Control
- Keep packages with desired version in internal/tools/go.mod
- Keep packages without version in bin/install_tools 
- - Call internal/tools/go.mod for using specific binary
- Specify dependencies in .github/dependabot.yml

# Practice
1. go get golang.org/x/tools/cmd/stringer 
- See the result with git diff (Indirect Dependencies)
2. We need to add golang.org/x/tools/cmd/stringer in tools.go for direct dependency
3. go mod tidy
4. git diff (Direct Dependencies)

# Resources
- [YouTube](https://www.youtube.com/watch?v=g_5n0W27XcY&list=PL7yAAGMOat_F7bOImcjx4ZnCtfyNEqzCy&index=5)
- [Repo (Same as Tutorial)](https://github.com/MarioCarrion/todo-api-microservice-example/tree/e44dc3d4016b80b8102eb1214fea452aebdd2667)
- [Repo (Most Recent)](https://github.com/MarioCarrion/todo-api-microservice-example)