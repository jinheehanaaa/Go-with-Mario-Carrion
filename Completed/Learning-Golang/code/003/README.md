# Module
- Module is a collection of Go packages stored in a file tree with a go.mod file at its root
- go.mod
- go.sum

# Versioning Module
- Module path
- go get github.com/jackc/pgx/v4@<Commit-Hash> for desired Commit
- go get github.com/jackc/pgx/v4@v4.10.1 for desired Version (Specific)
- go get -u github.com/jackc/pgx/v4 for updating package
- go list -m -u all for list & update available

# Resources
