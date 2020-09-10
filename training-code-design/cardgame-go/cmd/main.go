package main

import (
	"fmt"
	"os"

	"github.com/masu-mi/playground/training-code-design/cardgame-go"
	"github.com/masu-mi/playground/training-code-design/cardgame-go/blackjack"
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
		Use:   "blackjack",
		Short: "Example short usage",
		Args:  cobra.MaximumNArgs(3),
		RunE: func(_ *cobra.Command, args []string) error {
			fmt.Println("[Start] Subcommand head")
			fmt.Println("Name?")
			var name string
			fmt.Scan(&name)
			rs := cardgame.NewRounds(blackjack.NewGame, []string{name})
			rs.PlayAllRound()
			var result string
			switch rs.Players[0].Score {
			case 1:
				result = "Win"
			case 0:
				result = "Draw"
			default:
				result = "Loose"
			}
			fmt.Printf("%s: %s\n", name, result)
			return nil
		},
	}
	return cmd
}
