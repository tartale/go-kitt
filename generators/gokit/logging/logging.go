package logging

import (
	"os"
	"path"
	"text/template"

	"github.com/MarcGrol/golangAnnotations/model"

	"github.com/tartale/go-kitt/generators"
)

func Generate(parsedSourceData generators.ParsedSourceData) error {

	thisDir, err := generators.ThisDir()
	if err != nil {
		return err
	}
	tmpl := template.New("logging.tmpl").Funcs(generators.TemplateHelpers())
	tmpl = template.Must(tmpl.ParseFiles(path.Join(thisDir, "logging.tmpl")))

	for _, key := range parsedSourceData.Keys {
		parsedSource := parsedSourceData.Map[key]
		for _, intf := range parsedSource.Interfaces {
			data := struct {
				Interface model.Interface
			}{
				Interface: intf,
			}
			_ = tmpl.Execute(os.Stdout, data)
		}
	}

	return nil
}
