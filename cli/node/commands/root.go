package commands

import (
	"github.com/spf13/cobra"
)

func newRootCommand() *cobra.Command {
	rootCommand := &cobra.Command{
		Use:   "gotezos",
		Short: "Go Tezos is an implementation of a tezos node.",
	}

	rootCommand.AddCommand(
		newStartCommand(),
	)

	return rootCommand
}

func Execute() error {
	return newRootCommand().Execute()
}
