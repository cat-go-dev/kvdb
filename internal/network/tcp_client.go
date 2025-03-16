package network

import (
	"context"
	"fmt"
	"log/slog"
	"net"
)

type TCPClient struct {
	serverAddr string
	logger     *slog.Logger
}

func NewTCPClient(serverAddr string, logger *slog.Logger) *TCPClient {
	return &TCPClient{
		serverAddr: serverAddr,
		logger:     logger,
	}
}

const chanBuffSize = 10

func (c TCPClient) Start(ctx context.Context) (chan<- string, error) {
	logAttrs := []any{
		slog.String("component", "tcp_client"),
		slog.String("method", "Start"),
		slog.String("server_addr", c.serverAddr),
	}

	receiverChan := make(chan string, chanBuffSize)

	conn, err := net.Dial("tcp", c.serverAddr)
	if err != nil {
		wErr := fmt.Errorf("start tcp client: %w", err)
		c.logger.ErrorContext(ctx, wErr.Error(), logAttrs...)
		return receiverChan, fmt.Errorf("start tcp client: %w", err)
	}

	c.logger.InfoContext(ctx, "tcp server started", logAttrs...)

	go func() {
		defer conn.Close()

		for {
			select {
			case msg := <-receiverChan:
				c.logger.InfoContext(ctx, fmt.Sprintf("msg received: %s", msg), logAttrs...)
			case <-ctx.Done():
				c.logger.WarnContext(ctx, "canceled context", logAttrs...)
				return
			}
		}
	}()

	return receiverChan, nil
}
