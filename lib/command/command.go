package command

import (
	"os"
)

func Execute() {
	cmd := newCommand()

	cmd.SetOutput(os.Stdout)

	if err := cmd.Execute(); err != nil {
		cmd.SetOut(os.Stderr)

		cmd.PrintErr(err)

		os.Exit(1)
	}
}
