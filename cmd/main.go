package main

import (
	"github.com/pubgo/errors"
	"github.com/pubgo/go-readability/cmd/cmds"
	"os"
)

func main() {
	errors.ErrHandle(cmds.GoReadabilityCmd.Execute(), func(err *errors.Err) {
		err.P()
		os.Exit(-1)
	})
}
