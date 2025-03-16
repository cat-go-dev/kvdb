package network

import "errors"

var (
	errInvalidLogger      = errors.New("invalid logger")
	errEmptyServerAddress = errors.New("empty server address")
)
