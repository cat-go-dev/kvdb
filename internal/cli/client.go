package cli

import (
	"bufio"
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

type Executor interface {
	Execute(ctx context.Context, command string) string
}

type Client struct {
	e      Executor
	logger *slog.Logger
}

func NewClient(execuror Executor, logger *slog.Logger) (*Client, error) {
	if execuror == nil {
		return nil, errInvalidExecutor
	}

	if logger == nil {
		return nil, errInvalidLogger
	}

	return &Client{
		e:      execuror,
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
	commandCh := make(chan string)

	for {
		fmt.Print(commandPrefix)

		go func() {
			command, err := reader.ReadString('\n')
			if err != nil {
				wErr := fmt.Errorf("read command: %w", err)
				c.logger.ErrorContext(ctx, wErr.Error(), logAttrs...)
				fmt.Printf("%ssomething went wrong \r\n", commandPrefix)
				return
			}

			preparedCommand := strings.ReplaceAll(command, "\r\n", "")

			commandCh <- preparedCommand
		}()

		select {
		case command := <-commandCh:
			fmt.Printf("%s%s \r\n", commandPrefix, c.e.Execute(ctx, command))
		case <-ctx.Done():
			c.logger.WarnContext(ctx, "canceled context", logAttrs...)
			return errCanceledContext
		}
	}
}
