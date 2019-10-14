package gokit

import (
	"github.com/tartale/go-kitt/external/errorz"
	"github.com/tartale/go-kitt/generators"
	"github.com/tartale/go-kitt/generators/gokit/authz"
	"github.com/tartale/go-kitt/generators/gokit/endpoint"
	"github.com/tartale/go-kitt/generators/gokit/logger"
)

func Generate(parsedSources generators.ParsedSourceData) error {

	var errs errorz.Errors

	err := authz.Generate(parsedSources)
	if err != nil {
		errs = append(errs, err)
	}

	err = endpoint.Generate(parsedSources)
	if err != nil {
		errs = append(errs, err)
	}

	err = logger.Generate(parsedSources)
	if err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}
