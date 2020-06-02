package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/PuerkitoBio/goquery"
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
		Use: "htmlq",
	}

	cmd.AddCommand(
		newSubCommand(),
	)

	cmd.PersistentFlags().BoolVarP(&flagVerbose, "verbose", "v", false, "enable verbose log")

	return cmd
}

func newSubCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query",
		Short: "Example short usage",
		Args:  cobra.MaximumNArgs(3),
		RunE: func(_ *cobra.Command, args []string) error {
			for _, urlStr := range args {
				fmt.Printf("[INFO] target: %s\n", urlStr)
				b, _ := url.Parse(urlStr)
				doc, err := goquery.NewDocument(urlStr)
				if err != nil {
					fmt.Println(err)
					continue
				}
				doc.Find("a").Each(func(index int, s *goquery.Selection) {
					l, ok := s.Attr("href")
					if ok {
						relURL, _ := url.Parse(l)
						linkURL := b.ResolveReference(relURL)
						fmt.Printf("%5d: %s\n", index, linkURL.String())
					}
				})
			}
			return nil
		},
	}
	return cmd
}
