
# Go-KITT

### Go-Kit Two Thousand

Generates go-kit boilerplate code from annotations provided the comments of the business logic.

## Getting Started

### Pre-requisites

* go 1.12 or higher
* `$GOBIN` is in your `$PATH`

### Development

`go get github.com/tartale/go-kitt`

`go generate ./...`

`go build`

### Updating Dependencies

* Add a library dependency
  * `go mod download "github.com/org/repo"`
* Add a tool dependency
  * Add an import to `tools/tools.go` (it won't compile; ignore)
  * Add an install line to `tools/tools.go`

### TODO

* Coming Soon!
