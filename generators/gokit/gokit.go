package gokit

import (
	"github.com/tartale/go-kitt/generators"
	"github.com/tartale/go-kitt/generators/gokit/authz"
	"github.com/tartale/go-kitt/generators/gokit/endpoint"
	"github.com/tartale/go-kitt/generators/gokit/http"
	"github.com/tartale/go-kitt/generators/gokit/logging"
	"github.com/tartale/go-kitt/lib/errorz"
)

func Generate(parsedSources generators.ParsedSourceData) error {

	var errs errorz.Errors

	err := authz.Generate(parsedSources)
	if err != nil {
		errs = append(errs, err)
	}

	err = http.Generate(parsedSources)
	if err != nil {
		errs = append(errs, err)
	}

	err = endpoint.Generate(parsedSources)
	if err != nil {
		errs = append(errs, err)
	}

	err = logging.Generate(parsedSources)
	if err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}
