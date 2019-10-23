
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
* Add/remove library dependencies based on imports
  * `go mod tidy`
* Add a tool dependency
  * Add an import to `tools/tools.go` (this file is ignored in compilation)
  * Add a `go:generate go install ...` line to `tools/tools.go`


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
##### Template functions:
* Prefer "pipe"-style template functions.
* Before adding a new function, ensure it hasn't already been implemented in the included libraries:
  * [github.com/Masterminds/sprig](https://github.com/Masterminds/sprig)
  * [github.com/leekchan/gtf](https://github.com/leekchan/gtf)
* "Pipe"-style functions should be registered starting with a lower-case letter;
regular functions should be registered starting with an upper-case letter.
* Functions should try to return a reasonable default value when unexpected input is
given (for a pipe function, a reasonable default could be the unchanged input). This
is in contrast to the common practice of returning an error, which
would cause readability issues in a template.  Using `panic` could
also be useful to early-terminate the code generation
* While not a good general practice, use of `interface{}` parameter/return types and `variadic`
arguments in a template function can help increase the readability of
the template; one reason being that templates don't require casting.
The trade-off is with type-safety and clarity in the template function itself,
however, those issues are more easily surfaced in templates and code
generators than service-like programs.

### TODO (prioritized)

* Implement HTTP transport generation _**(in-progress)**_
  * Figure out how much the golangAnnotations library can help
  * Use method annotations to decode http.Request into endpoint request
  * Use method annotations to connect handlers to REST urls
* Implement more go-kit middleware
  * Figure out how much the golangAnnotations, gowrap, and kitgen libraries can help
  * Authorization
  * Tracing
* Implement middleware connectors
* Implement alternate transport(s)
  * gRPC
* Add support for plugins
* Reduce boilerplate for generators
* Implement alternate framework(s)
  * gRPC Ecosystem
* Add version stamp to generated code files
