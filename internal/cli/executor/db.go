package executor

import (
	"context"
	"fmt"
	"kdb/internal/ports"
	"log/slog"
)

type DB struct {
	db     ports.Database
	logger *slog.Logger
}

func NewDBExecutor(db ports.Database, logger *slog.Logger) (*DB, error) {
	if db == nil {
		return nil, errInvalidDatabase
	}

	if logger == nil {
		return nil, errInvalidLogger
	}

	return &DB{
		db:     db,
		logger: logger,
	}, nil
}

func (d DB) Execute(ctx context.Context, command string) string {
	logAttrs := []any{
		slog.String("component", "cli"),
		slog.String("method", "executeCimmand"),
	}

	r, err := d.db.Execute(ctx, command)
	if err != nil {
		wErr := fmt.Errorf("db executing: %w", err)
		d.logger.ErrorContext(ctx, wErr.Error(), logAttrs...)
		return err.Error()
	}

	return r.Msg
}
