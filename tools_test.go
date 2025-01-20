package toolkit_test

import (
	"testing"

	"github.com/kovalevsky0/go-toolkit-module"
)

func TestToolsRandomString(t *testing.T) {
	var tools toolkit.Tools

	s := tools.RandomString(8)
	got := len(s)

	if got != 8 {
		t.Errorf("Expected 8 random symbols but got %d", got)
	}
}
