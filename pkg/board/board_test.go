package newgameoflife

import (
	"testing"
)

func TestNewGameOfLife(t *testing.T) {
	t.Run("should be able to return something", func(t *testing.T) {
		board := NewGameOfLife(10, 10)
		if board == nil {
			t.Error("NewGameOfLife(10, 10) should not return nil")
		}
	})

}
