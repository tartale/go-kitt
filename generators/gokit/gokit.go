package gokit

import (
	"github.com/tartale/go/pkg/errors"

	"github.com/tartale/go-kitt/generators"
	"github.com/tartale/go-kitt/generators/gokit/endpoint"
	"github.com/tartale/go-kitt/generators/gokit/http"
	"github.com/tartale/go-kitt/generators/gokit/logging"
)

func Generate(parsedSources generators.ParsedSourceData) error {

	var errs errors.Errors

	err := casbin.Generate(parsedSources)
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
