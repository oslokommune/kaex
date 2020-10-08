package main

import (
	"bytes"
	"fmt"
	"github.com/oslokommune/kaex/pkg/api"
	"github.com/spf13/cobra"
)

var templateName string

const templateNameHelp = `Define what template to use. Kaex will first look in ~/.config/kaex/templates, then in
https://github.com/oslokommune/kaex/tree/master/templates
`

const initializeHelpLong = `scaffolds a template. By default, kaex init will scaffold an application.yaml`

func buildInitializeCommand(kaex api.Kaex) *cobra.Command {
	cmd := &cobra.Command{
		Use: "initialize",
		Aliases: []string{"init", "i"},
		Short: "scaffolds a template application.yaml",
		Long: `scaffolds a template application.yaml.`,
		RunE: func(_ *cobra.Command, args []string) error {
			var buffer bytes.Buffer

			err := api.FetchTemplate(kaex, &buffer, templateName)
			if err != nil {
				return fmt.Errorf("error fetching template: %w", err)
			}

			_, err = buffer.WriteTo(kaex.Out)
			if err != nil {
				return fmt.Errorf("error writing to output stream: %w", err)
			}

			return nil
		},
	}

	flags := cmd.Flags()

	flags.StringVarP(&templateName, "template", "t", "application.yaml", templateNameHelp)

	return cmd
}
