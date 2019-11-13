package casbin

import (
	"context"
	"errors"
)

type GetSubjectFunc func(ctx context.Context) (string, error)

var (
	ErrUnauthorized = errors.New("Unauthorized access")
)
