package main

import (
	"github.com/pubgo/errors"
	"github.com/pubgo/go-readability/cmd/cmds"
	"os"
)

func main() {
	rootCmd := cmds.RootCmd
	rootCmd.AddCommand(
		cmds.GoReadabilityCmd,
	)

	errors.ErrHandle(rootCmd.Execute(), func(err *errors.Err) {
		err.P()
		os.Exit(-1)
	})
}
