package database

import (
	"context"
	"fmt"
	"kdb/internal/database/compute"
	"log/slog"
	"time"
)

type Database struct {
	logger *slog.Logger
}

type Compute interface {
	ReadCommand() (*compute.Command, error)
}

func NewDatabase(logger *slog.Logger) *Database {
	return &Database{
		logger: logger,
	}
}

type Result struct {
	Msg string
}

func (d Database) Execute(ctx context.Context, commandStr string) (*Result, error) {
	if time.Now().UnixNano()%2 == 0 {
		return &Result{
			Msg: "success",
		}, nil
	} else {
		return nil, fmt.Errorf("test")
	}
	// todo: parse command

	// todo: execute command

	return nil, nil
}
