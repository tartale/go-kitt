package logging

import (
	"path"
	"text/template"

	"github.com/MarcGrol/golangAnnotations/model"
	"github.com/tartale/go/pkg/errors"

	"github.com/tartale/go-kitt/generators"
	"github.com/tartale/go-kitt/helpers"
)

func Generate(parsedSourceData generators.ParsedSourceData) error {

	thisDir, err := helpers.ThisDir()
	if err != nil {
		return err
	}
	tmpl := template.New("logging.tmpl").Funcs(helpers.TemplateHelpers())
	tmpl = template.Must(tmpl.ParseGlob(path.Join(thisDir, "logging*tmpl")))

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
			generatedFilePath := generators.NewGeneratedPathForInterface(intf, "Logging")
			err := generators.GenerateFile(generatedFilePath, tmpl, data)
			if err != nil {
				errs = append(errs, err)
				continue
			}
			generatedPaths = append(generatedPaths, generatedFilePath)
		}
	}

	if len(errs) > 0 {
		err := generatedPaths.RemoveAll()
		if err != nil {
			errs = append(errs, err)
		}
		return errs
	}

	err = generatedPaths.FinalizeAll()
	if err != nil {
		return err
	}

	return nil
}
