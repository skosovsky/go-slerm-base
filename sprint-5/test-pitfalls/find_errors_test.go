package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// go test ./...
// ok      github.com/skosovsky/go-slerm-base/sprint-5/test-pitfalls       0.177s

func TestFindErrorsWithIO(t *testing.T) {
	t.Run("Test find errors", func(t *testing.T) {
		result, err := FindErrorsIn("sample.txt")
		if err != nil {
			t.Errorf("FindErrorsIn() error: %v", err)
		}

		require.Equal(t, []string{"some error", "another error"}, result)
	})
}

func TestFindErrorsWithoutIO(t *testing.T) {
	t.Run("Test find errors", func(t *testing.T) {
		reader := strings.NewReader("one\ntwo\nsome error\nthree")
		result, err := FindErrorsInWithoutIO(reader)
		if err != nil {
			t.Errorf("FindErrorsInWithoutIO() error: %v", err)
		}

		require.Equal(t, []string{"some error"}, result)
	})
}
