package cli

import "errors"

var errInvalidDatabase = errors.New("invalid database")
var errInvalidLogger = errors.New("invalid logger")
var errCanceledContext = errors.New("canceled context")
