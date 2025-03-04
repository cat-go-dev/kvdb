package database

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"kdb/internal/database/compute"
	"kdb/internal/ports"
)

type Database struct {
	logger *slog.Logger
}

type Compute interface {
	ReadCommand() (*compute.Command, error)
}

func NewDatabase(logger *slog.Logger) (*Database, error) {
	if logger == nil {
		return nil, errInvalidLogger
	}

	return &Database{
		logger: logger,
	}, nil
}

func (d Database) Execute(ctx context.Context, commandStr string) (*ports.Result, error) {
	if time.Now().UnixNano()%2 == 0 {
		return &ports.Result{
			Msg: "success",
		}, nil
	} else {
		return nil, fmt.Errorf("test")
	}
	// todo: parse command

	// todo: execute command

	return nil, nil
}
