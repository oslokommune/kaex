package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"time"
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
			if err := fetchMinimalExample(); err != nil {
				return err
			}

			return nil
		},
	}
)

func init() {
	initCmd.Flags().BoolVarP(&fullExample, "full", "f", false, "use full template to scaffold rather than the minimal template")

	rootCmd.AddCommand(initCmd)
}

func fetchMinimalExample() error {
	var url string
	
	if fullExample {
		url = "https://raw.githubusercontent.com/deifyed/kaex/master/examples/application-full.yaml"
	} else {
		url = "https://raw.githubusercontent.com/deifyed/kaex/master/examples/application-minimal.yaml"
	}

	spaceClient := http.Client{
		Timeout: 2 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		return getErr
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return readErr
	}

	fmt.Printf("%s\n", body)
	
	return nil
}
