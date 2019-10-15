package endpoint

import (
	"os"
	"path"
	"text/template"

	"github.com/MarcGrol/golangAnnotations/model"

	"github.com/tartale/go-kitt/external/errorz"
	"github.com/tartale/go-kitt/generators"
)

func Generate(parsedSourceData generators.ParsedSourceData) error {

	thisDir, err := generators.ThisDir()
	if err != nil {
		return err
	}
	tmpl := template.New("endpoint.tmpl").Funcs(generators.TemplateHelpers())
	tmpl = template.Must(tmpl.ParseGlob(path.Join(thisDir, "endpoint*tmpl")))

	var errs errorz.Errors
	for _, key := range parsedSourceData.Keys {
		parsedSource := parsedSourceData.Map[key]
		for _, intf := range parsedSource.Interfaces {
			data := struct {
				Interface model.Interface
			}{
				Interface: intf,
			}
			err := tmpl.Execute(os.Stdout, data)
			if err != nil {
				errs = append(errs, err)
			}
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}
