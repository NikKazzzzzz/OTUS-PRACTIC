package timeghecker

import (
	"testing"
)

func TestGetCurrentTime(t *testing.T) {
	_, err := GetCurrentTime("time.google.com")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	_, err = GetCurrentTime("invalid.server")
	if err == nil {
		t.Error("Expected an error, but got none")
	}
}
