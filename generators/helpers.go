package generators

import (
	"path"
	"runtime"
	"text/template"

	"github.com/MarcGrol/golangAnnotations/model"
	"github.com/pkg/errors"
)

func ThisDir() (string, error) {
	if _, filename, _, ok := runtime.Caller(1); ok {
		return path.Dir(filename), nil
	}

	return "", errors.New("Unable to determine source file of caller")
}

func TemplateHelpers() template.FuncMap {
	return template.FuncMap{
		"hasGoKitAnnotation": func(i model.Interface) bool {
			_, ok := AnnotationRegistry.ResolveAnnotationByName(i.DocLines, GoKit)
			return ok
		},
	}
}
