package authz

import (
	"path"
	"text/template"

	"github.com/MarcGrol/golangAnnotations/model"
	"github.com/tartale/go/pkg/errors"

	"github.com/tartale/go-kitt/generators"
)

func Generate(parsedSourceData generators.ParsedSourceData) error {

	thisDir, err := generators.ThisDir()
	if err != nil {
		return err
	}
	tmpl := template.New("casbin.tmpl").Funcs(generators.TemplateHelpers())
	tmpl = template.Must(tmpl.ParseGlob(path.Join(thisDir, "casbin*tmpl")))

	var errs errors.Errors
	var generatedPaths generators.GeneratedPaths
	for _, key := range parsedSourceData.Keys {
		parsedSource := parsedSourceData.Map[key]
		for _, intf := range parsedSource.Interfaces {
			data := struct {
				Interface model.Interface
			}{
				Interface: intf,
			}
			generatedFilePath := generators.NewGeneratedPathForInterface(intf, "AuthzCasbin")
			err := generators.GenerateFile(generatedFilePath, tmpl, data)
			if err != nil {
				errs = append(errs, err)
				continue
			}
			generatedPaths = append(generatedPaths, generatedFilePath)
		}
	}

	if len(errs) > 0 {
		generatedPaths.RemoveAll()
		return errs
	}

	err = generatedPaths.FinalizeAll()
	if err != nil {
		return err
	}

	return nil
}
