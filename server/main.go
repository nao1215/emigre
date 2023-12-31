// Package main is emigre command entrypoint.
package main

import (
	"os"

	"github.com/nao1215/emigre/server/cmd"
)

// osExit is wrapper for  os.Exit(). It's for unit test.
var osExit = os.Exit //nolint

func main() {
	osExit(cmd.Execute())
}
