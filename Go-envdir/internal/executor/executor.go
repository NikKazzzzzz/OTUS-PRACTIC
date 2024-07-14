package executor

import (
	"errors"
	"os"
	"os/exec"
)

// RunCmd выполняет команду с указанными переменными окружения.
// Принимает срез строк cmd, представляющий команду и её аргументы, и карту env с переменными окружения.
// Возвращает код выхода команды.
func RunCmd(cmd []string, env map[string]string) int {
	command := exec.Command(cmd[0], cmd[1:]...)
	command.Env = append(os.Environ(), mapToEnv(env)...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Stdin = os.Stdin

	err := command.Run()
	if err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			return exitError.ExitCode()
		}
		return 1
	}
	return 0
}

func mapToEnv(env map[string]string) []string {
	var envVars []string
	for key, value := range env {
		envVars = append(envVars, key+"="+value)
	}
	return envVars
}
