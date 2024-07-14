package envdir

import (
	"os"
	"path/filepath"
	"strings"
)

// ReadDir считывает указанную директорию и возвращает карту переменных окружения
// в формате map[string]string, где ключи - имена файлов, а значения - их содержимое.
func ReadDir(dir string) (map[string]string, error) {
	envs := make(map[string]string)

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() {
			name := file.Name()
			context, err := os.ReadFile(filepath.Join(dir, name))
			if err != nil {
				return nil, err
			}
			envs[name] = strings.TrimSpace(string(context))
		}
	}
	return envs, nil
}
