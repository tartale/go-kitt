package logger

import (
	"os"
	"path"
	"text/template"

	"github.com/tartale/go-kitt/generators"
)

func Generate(parsedSources generators.ParsedSourceData) error {

	thisDir, err := generators.ThisDir()
	if err != nil {
		return err
	}
	tmpl := template.New("logger.tmpl").
		Funcs(generators.TemplateHelpers())
	tmpl, err = tmpl.ParseFiles(path.Join(thisDir, "logger.tmpl"))
	if err != nil {
		return err
	}

	for _, key := range parsedSources.Keys {
		parsedSource := parsedSources.Map[key]
		tmpl.Execute(os.Stdout, parsedSource)
	}

	return nil
}
