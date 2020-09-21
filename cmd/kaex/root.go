package main

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use: "kaex",
		Short: "Kubernetes Application.yaml EXpander",
		Long: "A tool for converting a simpler application.yaml into Kubernetes resources",
	}
)

func Execute() error {
	return rootCmd.Execute()
}
