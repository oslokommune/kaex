package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/oslokommune/kaex/pkg/api"
	"github.com/spf13/cobra"
)

var (
	save bool
	podonly bool
	expandCmd = &cobra.Command{
		Use: "expand",
		Aliases: []string{"exp", "x"},
		Short: "expands an application.yaml from stdin (default)",
		Long: "expands an application.yaml from stdin",
		RunE: func(_ *cobra.Command, args []string) error {
			var kubernetesResourceBuffer bytes.Buffer

			input, err := readStdin()
			if err != nil {
				return err
			}
			
			app, err := api.ParseApplication(input)
			if err != nil {
				return err
			}
			
			err = api.Expand(&kubernetesResourceBuffer, app, podonly)
			if err != nil {
				return err
			}
			
			if save == false {
				fmt.Println(kubernetesResourceBuffer.String())
				
				return nil
			}
			
			err = writeToFile(app.Name + ".yaml", kubernetesResourceBuffer)
			if err != nil {
				return err
			}

			return nil
		},
	}
)

func init() {
	expandCmd.Flags().BoolVarP(&save, "save", "s", false, "save the expanded Kubernetes resources to files")
	expandCmd.Flags().BoolVarP(&podonly, "pod-only", "p", false, "create a pod resource instead of a deployment")

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

func writeToFile(path string, buffer bytes.Buffer) error {
	err := ioutil.WriteFile(path, buffer.Bytes(), 0644)
	
	if err != nil {
		return err
	}
	
	return nil
}