package database

import (
	"context"
	"fmt"
	"log/slog"

	"kdb/internal/database/compute"
	"kdb/internal/ports"
)

type Database struct {
	compute *compute.Compute
	logger  *slog.Logger
}

func NewDatabase(compute *compute.Compute, logger *slog.Logger) (*Database, error) {
	if compute == nil {
		return nil, errInvalidCompute
	}

	if logger == nil {
		return nil, errInvalidLogger
	}

	return &Database{
		compute: compute,
		logger:  logger,
	}, nil
}

func (d Database) Execute(ctx context.Context, commandStr string) (*ports.Result, error) {
	logAttrs := []any{
		slog.String("component", "database"),
		slog.String("method", "Execute"),
	}

	command, err := d.compute.Parse(ctx, commandStr)
	if err != nil {
		wErr := fmt.Errorf("compute parse: %w", err)
		d.logger.ErrorContext(ctx, wErr.Error(), logAttrs...)
		return nil, err
	}

	fmt.Println("command", command)

	// todo: execute command

	return &ports.Result{
		Msg: "success",
	}, nil
}
