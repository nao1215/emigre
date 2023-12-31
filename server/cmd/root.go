// Package cmd is a package that contains subcommands for the emigre CLI command.
package cmd

import (
	"github.com/charmbracelet/log"
	"github.com/nao1215/emigre/server/api"
	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "emigre",
		Short: "emigre is an API server for an image sharing service",
		RunE: func(_ *cobra.Command, _ []string) error {
			return api.Run()
		},
	}
	cmd.CompletionOptions.DisableDefaultCmd = true
	cmd.SilenceUsage = true
	cmd.SilenceErrors = true
	cmd.AddCommand(newVersionCmd())
	cmd.AddCommand(newBugReportCmd())
	return cmd
}

// Execute run process.
func Execute() int {
	rootCmd := newRootCmd()

	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		return 1
	}
	return 0
}
