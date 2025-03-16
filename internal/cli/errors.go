package cli

import "errors"

var (
	errInvalidExecutor = errors.New("invalid executor")
	errInvalidLogger   = errors.New("invalid logger")
	errCanceledContext = errors.New("canceled context")
)
