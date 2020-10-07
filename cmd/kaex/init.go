package main

import (
	"bytes"
	"fmt"
	"github.com/oslokommune/kaex/pkg/api"
	"github.com/spf13/cobra"
	"os"
)

type InitializeOptions struct {
	FullExample bool
	Debugger bool
}

var (
	options = &InitializeOptions{}
	initCmd = &cobra.Command{
		Use: "initialize",
		Aliases: []string{"init", "i"},
		Short: "scaffolds a template application.yaml",
		Long: `scaffolds a template application.yaml. Use --full to enhance the template with more than only the required settings.
		`,
		RunE: func(_ *cobra.Command, args []string) error {
			var err error

			var buffer bytes.Buffer

			if options.Debugger {
				err = api.GenerateDebugger(&buffer)

				if err != nil {
					fmt.Fprint(os.Stderr, err)
				} else {
					fmt.Fprint(os.Stdout, buffer.String())
				}

				return nil
			}

			if options.FullExample {
				err = api.FetchFullExample(&buffer)
			} else {
				err = api.FetchMinimalExample(&buffer)
			}

			if err != nil {
				return err
			}

			fmt.Fprintln(os.Stdout, buffer.String())

			return nil
		},
	}
)

func init() {
	initCmd.Flags().BoolVarP(&options.FullExample, "full", "f", false, "use full template to scaffold rather than the minimal template")
	initCmd.Flags().BoolVarP(&options.Debugger, "debugger", "d", false, "scaffold a debugger instance")

	rootCmd.AddCommand(initCmd)
}
