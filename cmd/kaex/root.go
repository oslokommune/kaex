package main

import (
	"github.com/oslokommune/kaex/pkg/api"
	"github.com/spf13/cobra"
	"os"
	"path"
)

var (
	rootCmd = &cobra.Command{
		Use: "kaex",
		Short: "Kubernetes Application.yaml EXpander",
		Long: "A tool for converting a simpler application.yaml into Kubernetes resources",
	}
)

func Execute() error {
	kaex := api.Kaex{
		Err: 		rootCmd.ErrOrStderr(),
		Out:        rootCmd.OutOrStdout(),
		In:         rootCmd.InOrStdin(),

		RepoURL:    "https://raw.githubusercontent.com/oslokommune/kaex/master",
	}

	configPath, err := os.UserConfigDir()
	if err == nil {
		kaex.ConfigPath = path.Join(configPath, "kaex")
	}

	rootCmd.AddCommand(buildInitializeCommand(kaex))
	rootCmd.AddCommand(buildExpandCommand(kaex))

	return rootCmd.Execute()
}
