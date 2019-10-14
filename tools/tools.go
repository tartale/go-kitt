// +build tools
package tools

import (
    _ "github.com/nyarly/inlinefiles"
    _ "github.com/go-kit/kit/cmd/kitgen"
)

//go:generate go install github.com/go-kit/kit/cmd/kitgen
//go:generate go get github.com/MarcGrol/golangAnnotations
//go:generate go generate github.com/MarcGrol/golangAnnotations
//go:generate go install github.com/MarcGrol/golangAnnotations
