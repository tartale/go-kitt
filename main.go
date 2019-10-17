package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/MarcGrol/golangAnnotations/model"
	"github.com/MarcGrol/golangAnnotations/parser"

	"github.com/tartale/go-kitt/generators"
	"github.com/tartale/go-kitt/generators/gokit"
	"github.com/tartale/go-kitt/lib/errorz"
)

func main() {
	var inputPath string
	if len(os.Args) == 1 {
		cwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		inputPath = cwd
	} else if len(os.Args) == 2 {
		inputPath = os.Args[1]
	} else {
		panic(fmt.Sprintf("usage: %s [<path>]", os.Args[0]))
	}
	inputPath = "/Users/tom.artale/Projects/go-beverage/domain"

	parsedSourceMap := parseAll(inputPath)
	parsedSources := generators.ParsedSourceData{
		Map:  parsedSourceMap,
		Keys: parsedSourceMap.Keys(),
	}

	var errs errorz.Errors

	err := gokit.Generate(parsedSources)
	if err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		panic(errs)
	}
}

func parseAll(inputPath string) generators.ParsedSourceMap {

	const (
		excludeMatchPattern = "^.*" + generators.GeneratedPathNewExtension + "$" + "|" +
			"^.*" + generators.GeneratedPathExtension + "$"
		includeMatchPattern = "^.*.go$"
	)

	result := make(generators.ParsedSourceMap)

	info, err := os.Stat(inputPath)
	if err != nil {
		panic(err)
	}

	if !info.IsDir() {
		inputPath = path.Dir(inputPath)
	}

	err = filepath.Walk(inputPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			return nil
		}
		parsedSources, err := parser.New().ParseSourceDir(path, includeMatchPattern, excludeMatchPattern)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return nil
		}
		if !isEmpty(parsedSources) {
			result[path] = parsedSources
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	return result
}

func isEmpty(parsedSources model.ParsedSources) bool {
	return len(parsedSources.Operations) == 0 &&
		len(parsedSources.Interfaces) == 0 &&
		len(parsedSources.Typedefs) == 0 &&
		len(parsedSources.Structs) == 0 &&
		len(parsedSources.Enums) == 0
}
