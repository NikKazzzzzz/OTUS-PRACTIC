package main

import (
	"fmt"
	"github.com/NikKazzzzzz/OTUS-PRACTIC/Go-envdir/internal/envdir"
	"github.com/NikKazzzzzz/OTUS-PRACTIC/Go-envdir/internal/executor"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, `Usage: Go-envdir /path/to/envdir command arg1 arg2...`)
		os.Exit(1)
	}

	envDir := os.Args[1]
	cmd := os.Args[2:]

	envs, err := envdir.ReadDir(envDir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading directory: ", err)
		os.Exit(1)
	}

	exitCode := executor.RunCmd(cmd, envs)
	os.Exit(exitCode)
}
