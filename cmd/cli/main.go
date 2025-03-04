package main

import (
	"context"
	"fmt"
	"kdb/internal/cli"
	"kdb/internal/database"
	"log/slog"
	"os"
)

func main() {
	ctx := context.Background()
	// todo: maybe simple logs (without JSON) for localhost
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	database := &database.Database{}

	client, err := cli.NewClient(database, logger)
	if err != nil {
		wErr := fmt.Errorf("creating new cli client: %w", err)
		logger.ErrorContext(ctx, wErr.Error())
		return
	}

	err = client.Run(ctx)
	if err != nil {
		wErr := fmt.Errorf("running cli client: %w", err)
		logger.ErrorContext(ctx, wErr.Error())
		return
	}
}
