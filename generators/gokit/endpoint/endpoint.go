package endpoint

import (
	"path"
	"text/template"

	"github.com/MarcGrol/golangAnnotations/model"

	"github.com/tartale/go-kitt/generators"
	"github.com/tartale/go-kitt/helpers"
	"github.com/tartale/go/pkg/errors"
)

func Generate(parsedSourceData generators.ParsedSourceData) error {

	thisDir, err := helpers.ThisDir()
	if err != nil {
		return err
	}
	tmpl := template.New("endpoint.tmpl").Funcs(helpers.TemplateHelpers())
	tmpl = template.Must(tmpl.ParseGlob(path.Join(thisDir, "endpoint*tmpl")))

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
			generatedFilePath := generators.NewGeneratedPathForInterface(intf, "Endpoint")
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
