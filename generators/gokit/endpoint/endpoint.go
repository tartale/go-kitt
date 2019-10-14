package endpoint

import (
    "fmt"
    "html/template"
    "path"

    "github.com/tartale/go-kitt/generators"
)

func Generate(parsedSources generators.ParsedSourceData) error {

	thisDir, err := generators.ThisDir()
	if err != nil {
		return err
	}
	tmpl := template.Must(template.ParseFiles(path.Join(thisDir, "endpoint.tmpl")))

	for _, key := range parsedSources.Keys {
	    parsedSource := parsedSources.Map[key]
	    fmt.Println(key)
	    fmt.Println(parsedSource)
        tmpl = tmpl
    }

	return nil
}
