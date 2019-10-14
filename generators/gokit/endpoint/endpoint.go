package endpoint

import (
	"path"
	"text/template"

	"github.com/tartale/go-kitt/generators"
)

func Generate(parsedSources generators.ParsedSourceData) error {

	thisDir, err := generators.ThisDir()
	if err != nil {
		return err
	}
	tmpl := template.Must(template.ParseFiles(path.Join(thisDir, "endpoint.tmpl"))).
		Funcs(generators.TemplateHelpers())

	for _, key := range parsedSources.Keys {
		parsedSource := parsedSources.Map[key]
		tmpl = tmpl
		parsedSource = parsedSource
	}

	return nil
}
