package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"sigs.k8s.io/yaml"

	"github.com/deifyed/kaex/pkg/api"
)

var (
	save bool
	expandCmd = &cobra.Command{
		Use: "expand",
		Aliases: []string{"exp", "x"},
		Short: "expands an application.yaml from stdin (default)",
		Long: "expands an application.yaml from stdin",
		RunE: func(_ *cobra.Command, args []string) error {
			if err := expand(); err != nil {
				return err
			}

			return nil
		},
	}
)

func init() {
	// Yet to be implemented
	//expandCmd.Flags().BoolVarP(&save, "save", "s", false, "save the expanded Kubernetes resources to files")
	save = false
	
	rootCmd.AddCommand(expandCmd)
}

func readStdin() (string, error) {
	lines := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()
		lines = append(lines, text)
	}
	
	return strings.Join(lines, "\n"), nil
}

func parseApplication(raw string) (api.Application, error) {
	var app api.Application

	err := yaml.Unmarshal([]byte(raw), &app)
	if err != nil {
		return api.Application{}, err
	}
	
	return app, nil
}

func writeResource(w io.Writer, resource interface{}) {
	fmt.Fprintf(w, "%s\n---", resource)
}

func expand() error {
	var buffer bytes.Buffer

	input, err := readStdin()
	if err != nil {
		return err
	}
	
	app, err := parseApplication(input)
	if err != nil {
		return err
	}

	service, err := yaml.Marshal(api.CreateService(app))
	if err != nil {
		return err
	}
	writeResource(&buffer, service)
	
	fmt.Print(buffer.String())

	return nil
}