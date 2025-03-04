package cli

import (
	"bufio"
	"context"
	"fmt"
	"kdb/internal/database"
	"log/slog"
	"os"
)

type Client struct {
	db     *database.Database
	logger *slog.Logger
}

func NewClient(db *database.Database, logger *slog.Logger) (*Client, error) {
	if db == nil {
		return nil, errInvalidDatabase
	}

	if logger == nil {
		return nil, errInvalidLogger
	}

	return &Client{
		db:     db,
		logger: logger,
	}, nil
}

const commandPrefix = "[kdb] > "

func (c Client) Run(ctx context.Context) error {
	logAttrs := []any{
		slog.String("component", "cli"),
		slog.String("method", "Run"),
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(commandPrefix)

		command, err := reader.ReadString('\n')
		if err != nil {
			wErr := fmt.Errorf("read command: %w", err)
			c.logger.ErrorContext(ctx, wErr.Error(), logAttrs...)
			fmt.Printf("%ssomething went wrong \r\n", commandPrefix)
			continue
		}

		fmt.Printf("%s%s \r\n", commandPrefix, c.executeCommand(ctx, command))
	}
}

func (c Client) executeCommand(ctx context.Context, command string) string {
	logAttrs := []any{
		slog.String("component", "cli"),
		slog.String("method", "executeCimmand"),
	}

	r, err := c.db.Execute(ctx, command)
	if err != nil {
		wErr := fmt.Errorf("db executing: %w", err)
		c.logger.ErrorContext(ctx, wErr.Error(), logAttrs...)
		return fmt.Sprintf("error while executing command: %s", command)
	}

	return r.Msg
}
