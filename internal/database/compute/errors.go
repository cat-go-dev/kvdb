package compute

import "errors"

var (
	errInvalidLogger      = errors.New("invalid logger")
	errNotEnoughArguments = errors.New("not enought arguments")
	errInvalidArgument    = errors.New("invalid argument")
	errUnknownCommandType = errors.New("unknown command type")
)
