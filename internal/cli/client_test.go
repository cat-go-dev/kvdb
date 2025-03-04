package cli

import (
	"kdb/internal/database"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClientWithEmptyDatabase(t *testing.T) {
	client, err := NewClient(nil, nil)

	assert.Nil(t, client)
	assert.ErrorIs(t, err, errInvalidDatabase)
}

func TestNewClientWithEmptyLogger(t *testing.T) {
	db := &database.Database{}
	client, err := NewClient(db, nil)

	assert.Nil(t, client)
	assert.ErrorIs(t, err, errInvalidLogger)
}

func TestExecuteCommand(t *testing.T) {
	logger := slog.New(&slog.TextHandler{})
	db := database.NewDatabase(logger)

}
