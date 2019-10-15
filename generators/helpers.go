package generators

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"text/template"

	"github.com/MarcGrol/golangAnnotations/model"
	"github.com/pkg/errors"

	"github.com/tartale/go-kitt/config"
)

func ThisDir() (string, error) {
	if _, filename, _, ok := runtime.Caller(1); ok {
		return path.Dir(filename), nil
	}

	return "", errors.New("Unable to determine source file of caller")
}

func TemplateHelpers() template.FuncMap {
	return template.FuncMap{
		"ShouldGenerateLogging": ShouldGenerateLogging,
		"MethodSignature":       MethodSignature,
		"FieldDeclaration":      FieldDeclaration,
		"FieldDeclarations":     FieldDeclarations,
		"FieldNames":            FieldNames,
	}
}

func DocLines(obj interface{}) []string {

	switch v := obj.(type) {
	case model.Struct:
		return v.DocLines
	case model.Operation:
		return v.DocLines
	case model.Interface:
		return v.DocLines
	case model.Typedef:
		return v.DocLines
	case model.Enum:
		return v.DocLines
	default:
		return []string{}
	}

	return []string{}
}

func ShouldGenerateLogging(objs ...interface{}) bool {
	result := config.Config.LogAllMethods
	for _, obj := range objs {
		_, ok := AnnotationRegistry.ResolveAnnotationByName(DocLines(obj), Logging)
		if ok {
			result = true
			break
		}
	}
	return result
}

func MethodSignature(obj model.Operation) string {
	return fmt.Sprintf("%s(%s) (%s)", obj.Name, FieldDeclarations(obj.InputArgs), FieldDeclarations(obj.OutputArgs))
}

func FieldDeclaration(field model.Field) string {
	var fieldModifier string
	if field.IsSlice {
		fieldModifier = fmt.Sprintf("%s%s", fieldModifier, "[]")
	}
	if field.IsPointer {
		fieldModifier = fmt.Sprintf("%s%s", fieldModifier, "*")
	}
	fieldType := fmt.Sprintf("%s%s", fieldModifier, field.TypeName)
	return fmt.Sprintf("%s %s", field.Name, fieldType)
}

func FieldDeclarations(fields []model.Field) string {
	var result []string
	for i, field := range fields {
		if field.Name == "" {
			field.Name = fmt.Sprintf("v%d", i)
		}
		result = append(result, FieldDeclaration(field))
	}

	return strings.Join(result, ", ")
}

func FieldNames(fields []model.Field) string {
	var result []string
	for i, field := range fields {
		if field.Name == "" {
			field.Name = fmt.Sprintf("v%d", i)
		}
		result = append(result, field.Name)
	}

	return strings.Join(result, ", ")
}
