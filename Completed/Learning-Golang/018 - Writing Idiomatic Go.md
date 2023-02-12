# 1. Use gofmt
- Use gofmt for consistency
- - for nvim user: <code>:GoInstallBinaries</code>

# 2. Keep receiver names short
- first character of type name
- avoid self, this, me
- 3rd party linter plugin: <code>golangci-lint run ./..</code> for detecting consistency
- - brew install golanggci-lint

# 3.Name packages accordingly
## Example: todo-microservice
### elasticsearch
- Elastic Search
### postgresql
- Accessing postgresql DB

# 4. Group imports by their origin

# 5. Use short names for variables with limited scope
q["from"] = args.From => query["from"] = args.From
