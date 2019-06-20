package main

import (
	"github.com/pubgo/assert"
	"github.com/pubgo/go-readability/cmd/cmds"
	"os"
)

func main() {
	rootCmd := cmds.RootCmd
	rootCmd.AddCommand(
		cmds.GoReadabilityCmd,
	)

	assert.ErrHandle(rootCmd.Execute(), func(err *assert.KErr) {
		err.P()
		os.Exit(-1)
	})
}
