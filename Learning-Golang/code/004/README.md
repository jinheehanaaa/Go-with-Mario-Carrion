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