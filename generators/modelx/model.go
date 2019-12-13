package modelx

import (
	"fmt"
	"strings"

	"github.com/MarcGrol/golangAnnotations/model"
)

type Declaration struct {
	Field model.Field
}

type Declarations []Declaration

func (o Declaration) String() string {

	result := o.Field.Name
	typeName := o.Field.TypeName
	if o.Field.IsSlice {
		typeName = "[]" + typeName
	}
	if o.Field.IsPointer {
		typeName = "*" + typeName
	}
	result = fmt.Sprintf("%s %s %s", result, typeName, o.Field.Tag)

	return result
}

type Assignment struct {
	Field model.Field
	RHS   string
}

type Assignments []Assignment

func (o Assignment) String() string {

	return fmt.Sprintf("%s := %s", o.Field.Name, o.RHS)
}

type AuthZEnforcement map[string]string

func (o AuthZEnforcement) Assignments() string {

	var results []string

	if obj, ok := o["obj"]; ok {
		results = append(results, fmt.Sprintf(`%s := "%s"`, "obj", obj))
	}
	if act, ok := o["act"]; ok {
		results = append(results, fmt.Sprintf(`%s := "%s"`, "act", act))
	}

	return strings.Join(results, "\n")
}

func (o AuthZEnforcement) Params() string {

	var results []string

	if _, ok := o["obj"]; ok {
		results = append(results, "obj")
	}
	if _, ok := o["act"]; ok {
		results = append(results, "act")
	}

	return strings.Join(results, ", ")
}
