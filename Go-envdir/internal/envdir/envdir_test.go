package envdir

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadDir(t *testing.T) {
	dir, err := os.MkdirTemp("", "envdir")
	if err != nil {
		t.Fatal(err)
	}
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {

		}
	}(dir)

	os.WriteFile(filepath.Join(dir, "VAR1"), []byte("value1"), 0644)
	os.WriteFile(filepath.Join(dir, "VAR2"), []byte("value2"), 0644)

	envs, err := ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}

	if envs["VAR1"] != "value1" {
		t.Errorf("want VAR1=%q, got %q", "value1", envs["VAR1"])
	}
	if envs["VAR2"] != "value2" {
		t.Errorf("want VAR2=%q, got %q", "value2", envs["VAR2"])
	}
}
