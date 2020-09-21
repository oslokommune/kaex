package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
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

func expand() error {
	input, err := readStdin()
	if err != nil {
		return err
	}
	
	fmt.Println(input)

	return nil
}