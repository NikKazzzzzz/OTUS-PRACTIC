package executor

import "testing"

func TestRunCmd(t *testing.T) {
	cmd := []string{"echo", "hello"}
	env := map[string]string{"GREETING": "hello"}

	exitCode := RunCmd(cmd, env)
	if exitCode != 0 {
		t.Errorf("exit code should be 0, but it is %d", exitCode)
	}
}

func TestRunCmdWithEnv(t *testing.T) {
	cmd := []string{"false"}
	env := map[string]string{}

	exitCode := RunCmd(cmd, env)
	if exitCode != 1 {
		t.Errorf("exit code should be 1, but it is %d", exitCode)
	}
}
