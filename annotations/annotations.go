package annotations

import (
	"sync"

	"github.com/MarcGrol/golangAnnotations/generator/annotation"
	"github.com/MarcGrol/golangAnnotations/generator/rest/restAnnotation"
)

const (
	TypeFramework      = "Framework"
	ParamFrameworkName = "name"
	TypeLogging        = "Logging"
	TypeAuthorization  = "Authorization"
	ParamEnforcement   = "enforcement"
)

func frameworkNames() map[string]bool {
	return map[string]bool{"gokit": true}
}

var descriptors []annotation.AnnotationDescriptor
var descriptorsInit sync.Once

func Descriptors() []annotation.AnnotationDescriptor {
	descriptorsInit.Do(func() {
		descriptors = append(descriptors, annotation.AnnotationDescriptor{
			Name:       TypeFramework,
			ParamNames: []string{ParamFrameworkName},
			Validator: func(annot annotation.Annotation) bool {
				v, hasParam := annot.Attributes[ParamFrameworkName]
				_, paramIsLegal := frameworkNames()[v]

				return annot.Name == TypeFramework && hasParam && paramIsLegal

			}})

		descriptors = append(descriptors, annotation.AnnotationDescriptor{
			Name:       TypeLogging,
			ParamNames: []string{},
			Validator: func(annot annotation.Annotation) bool {
				return annot.Name == TypeLogging &&
					len(annot.Attributes) == 0
			},
		})

		descriptors = append(descriptors, annotation.AnnotationDescriptor{
			Name:       TypeAuthorization,
			ParamNames: []string{ParamEnforcement},
			Validator: func(annot annotation.Annotation) bool {
				return annot.Name == TypeAuthorization &&
					len(annot.Attributes) == 1
			},
		})

		descriptors = append(descriptors, restAnnotation.Get()...)
	})

	return descriptors
}

func AddDescriptor(descriptor annotation.AnnotationDescriptor) {
	descriptors = append(descriptors, descriptor)
}

func AnnotationRegistry() annotation.AnnotationRegister {

	return annotation.NewRegistry(Descriptors())
}
