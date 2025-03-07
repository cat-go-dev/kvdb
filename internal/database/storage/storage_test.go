package storage

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	tests := []struct {
		name        string
		key         string
		expected    string
		expectedErr error
	}{
		{name: "success case", key: "test", expected: "val", expectedErr: nil},
		{name: "failed case", key: "test", expected: "", expectedErr: fmt.Errorf("some err")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}
