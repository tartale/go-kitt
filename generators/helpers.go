package generators

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"text/template"

	"github.com/MarcGrol/golangAnnotations/model"
	"github.com/Masterminds/sprig"
	"github.com/leekchan/gtf"
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
	result := template.FuncMap{
		"HeaderComment":         HeaderComment,
		"ShouldGenerateLogging": ShouldGenerateLogging,
		"MethodSignature":       MethodSignature,
		"FieldDeclaration":      FieldDeclaration,
		"FieldDeclarations":     FieldDeclarations,
		"FieldNames":            FieldNames,
		"FieldCtxFilter":        FieldCtxFilter,
	}
	for k, v := range sprig.GenericFuncMap() {
		result[k] = v
	}
	// hack until sprig@v3.x is available
	result["get"] = func(d map[string]interface{}, key string) interface{} {
		if val, ok := d[key]; ok {
			return val
		}
		return ""
	}
	for k, v := range gtf.GtfTextFuncMap {
		result[k] = v
	}

	return result
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

const headerComment = `
/**
 * DO NOT EDIT - changes will be over-written
 * Generated by github.com/tartale/go-kitt
**/
`

func HeaderComment() string {
	return headerComment
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
	return fmt.Sprintf("%s(%s) (%s)", obj.Name,
		strings.Join(FieldDeclarations(obj.InputArgs), ", "),
		strings.Join(FieldDeclarations(obj.OutputArgs), ", "))
}

func FieldDeclaration(field model.Field, index int) string {
	var fieldModifier string
	if field.IsSlice {
		fieldModifier = fmt.Sprintf("%s%s", fieldModifier, "[]")
	}
	if field.IsPointer {
		fieldModifier = fmt.Sprintf("%s%s", fieldModifier, "*")
	}
	fieldType := fmt.Sprintf("%s%s", fieldModifier, field.TypeName)
	if field.Name == "" {
		field.Name = fmt.Sprintf("v%d", index)
	}

	return fmt.Sprintf("%s %s", field.Name, fieldType)
}

func FieldDeclarations(fields []model.Field) []string {
	var result []string
	for i, field := range fields {
		result = append(result, FieldDeclaration(field, i))
	}

	return result
}

func FieldNames(fields []model.Field) []string {
	var result []string
	for i, field := range fields {
		if field.Name == "" {
			field.Name = fmt.Sprintf("v%d", i)
		}
		result = append(result, field.Name)
	}

	return result
}

func FieldCtxFilter(fields []model.Field) []model.Field {
	var result []model.Field
	for _, field := range fields {
		if field.TypeName == "context.Context" {
			continue
		}
		result = append(result, field)
	}
	return result
}
