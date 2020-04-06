package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sync"

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
		Use:  "translate",
		RunE: translate,
	}
	cmd.PersistentFlags().BoolVarP(&flagVerbose, "verbose", "v", false, "enable verbose log")
	return cmd
}

func translate(_ *cobra.Command, args []string) error {
	pr, pw := io.Pipe()
	tr := io.TeeReader(os.Stdin, pw)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		fmt.Println("START ioutil.ReadAll()")
		buf, _ := ioutil.ReadAll(pr)
		fmt.Println("GO FUNC: " + string(buf))
		wg.Done()
	}()
	fmt.Println("example")
	io.Copy(os.Stderr, tr)
	buf := bytes.NewBuffer([]byte{})
	buf := bytes.NewBufferString("")
	pw.Close()
	fmt.Println("end")
	wg.Wait()
	return nil
}
