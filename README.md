
# Go-KITT

### Go-Kit Two Thousand

Generates go-kit boilerplate code from annotations provided the comments of the business logic.

## Getting Started

### Pre-requisites

* go 1.12 or higher installed and in your `$PATH`
* `$GOBIN` is in your `$PATH`

### Development

`go get github.com/tartale/go-kitt`

`go generate ./...`

`go build`

### Updating Dependencies

* Add a library dependency
  * `go mod download "github.com/org/repo"`
* Add a tool dependency
  * Add an import to `tools/tools.go` (it won't compile; just ignore the errors)
  * Add an `//go:generate` install line to `tools/tools.go`

### Template Design Guidelines
* Prefer descriptive variables over `{{ . }}`-like references
* Try to declare variables at the top, instead of near the code where they're used.
This will help to keep the template readable.
* Try to design a single "main" template, and separate out "sub-templates"
for "one-shot" and "repeated" sections, for readability.
* Start all "instruction" lines with `{{-`, to prevent extra newlines
in the generated code
* Use indentations in "instruction" lines, and include
a "buffer space" within the double-brackets, to enhance readability.

### TODO (prioritized)

-[ ] Implement HTTP transport generation (in-progress)
   -[ ] Figure out how much the golangAnnotations library can help
   -[ ] Use method annotations to decode http.Request into endpoint request
   -[ ] Use method annotations to connect handlers to REST urls
-[ ] Implement more go-kit middleware
   -[ ] Figure out how much the golangAnnotations, gowrap, and kitgen libraries can help
   -[ ] Authorization
   -[ ] Tracing
-[ ] Implement middleware connectors
-[ ] Implement alternate transport(s)
   -[ ] gRPC
-[ ] Add support for plugins
-[ ] Reduce boilerplate for generators
-[ ] Implement alternate framework(s)
   -[ ] gRPC Ecosystem
-[ ] Add version stamp to generated code files
