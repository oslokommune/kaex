package main

import (
	"bytes"
	"fmt"
	"github.com/oslokommune/kaex/pkg/api"
	"github.com/spf13/cobra"
)

var (
	fullExample bool
	initCmd = &cobra.Command{
		Use: "initialize",
		Aliases: []string{"init", "i"},
		Short: "scaffolds a template application.yaml",
		Long: `scaffolds a template application.yaml. Use --full to enhance the template with more than only the required settings.
		`,
		RunE: func(_ *cobra.Command, args []string) error {
			var err error
			
			var buffer bytes.Buffer
			
			if fullExample {
				err = api.FetchFullExample(&buffer)
			} else {
				err = api.FetchMinimalExample(&buffer)
			}
			
			if err != nil {
				return err
			}
			
			fmt.Println(buffer.String())

			return nil
		},
	}
)

func init() {
	initCmd.Flags().BoolVarP(&fullExample, "full", "f", false, "use full template to scaffold rather than the minimal template")

	rootCmd.AddCommand(initCmd)
}
