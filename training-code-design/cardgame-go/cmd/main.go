package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	cmd := newFoobarCommand()
	return cmd.Execute()
}

func newFoobarCommand() *cobra.Command {
	var (
		flagVerbose bool
	)

	cmd := &cobra.Command{
		Use: "card",
	}

	cmd.AddCommand(
		newSubCommand(),
	)

	cmd.PersistentFlags().BoolVarP(&flagVerbose, "verbose", "v", false, "enable verbose log")

	return cmd
}

func newSubCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "head",
		Short: "Example short usage",
		Args:  cobra.MaximumNArgs(3),
		RunE: func(_ *cobra.Command, args []string) error {
			fmt.Println("[Start] Subcommand head")
			fmt.Printf("%v\n", args)
			//
			return nil
		},
	}
	return cmd
}
