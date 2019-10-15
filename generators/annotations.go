package generators

import "github.com/MarcGrol/golangAnnotations/generator/annotation"

const (
	GoKit   = "GoKit"
	Logging = "Logging"
)

var AnnotationRegistry = annotation.NewRegistry(
	[]annotation.AnnotationDescriptor{
		{
			Name:       GoKit,
			ParamNames: []string{},
			Validator: func(annot annotation.Annotation) bool {
				return annot.Name == GoKit &&
					len(annot.Attributes) == 0
			},
		},
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
