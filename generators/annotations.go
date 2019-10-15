package generators

import "github.com/MarcGrol/golangAnnotations/generator/annotation"

const (
	Logging = "Logging"
)

var AnnotationRegistry = annotation.NewRegistry(
	[]annotation.AnnotationDescriptor{
		{
			Name:       Logging,
			ParamNames: []string{},
			Validator: func(annot annotation.Annotation) bool {
				return annot.Name == Logging &&
					len(annot.Attributes) == 0
			},
		},
	},
)
