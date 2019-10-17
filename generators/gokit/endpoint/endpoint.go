package endpoint

import (
	"path"
	"text/template"

	"github.com/MarcGrol/golangAnnotations/model"

	"github.com/tartale/go-kitt/generators"
	"github.com/tartale/go-kitt/lib/errorz"
)

func Generate(parsedSourceData generators.ParsedSourceData) error {

	thisDir, err := generators.ThisDir()
	if err != nil {
		return err
	}
	tmpl := template.New("endpoint.tmpl").Funcs(generators.TemplateHelpers())
	tmpl = template.Must(tmpl.ParseGlob(path.Join(thisDir, "endpoint*tmpl")))

	var errs errorz.Errors
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
