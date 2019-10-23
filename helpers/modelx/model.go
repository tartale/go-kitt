package modelx

import (
	"fmt"

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
