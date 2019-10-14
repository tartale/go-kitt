package generators

import (
	"path"
	"runtime"

	"github.com/pkg/errors"
)

func ThisDir() (string, error) {
	if _, filename, _, ok := runtime.Caller(1); ok {
		return path.Dir(filename), nil
	}

	return "", errors.New("Unable to determine source file of caller")
}
