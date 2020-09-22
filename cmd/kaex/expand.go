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
	podonly bool
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
	expandCmd.Flags().BoolVarP(&podonly, "pod-only", "p", false, "create a pod resource instead of a deployment")
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

func writeResource(w io.Writer, resource interface{}) error {
	serializedResource, err := yaml.Marshal(resource)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(w, "%s\n---\n", serializedResource)
	if err != nil {
		return err
	}
	
	return nil
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

	if app.Port != 0 {
		service, err := api.CreateService(app)
		if err != nil {
			return err
		}
		err = writeResource(&buffer, service)
		if err != nil {
			return err
		}
	}

	if app.Url != "" {
		ingress, err := api.CreateIngress(app)
		if err != nil {
			return err
		}
		err = writeResource(&buffer, ingress)
		if err != nil {
			return err
		}
	}

	if podonly == false {
		deployment, err := api.CreateDeployment(app)
		if err != nil {
			return err
		}
		err = writeResource(&buffer, deployment)
		if err != nil {
			return err
		}
	} else {
		pod, err := api.CreatePod(app)
		if err != nil {
			return err
		}
		err = writeResource(&buffer, pod)
		if err != nil {
			return err
		}
	}

	fmt.Print(buffer.String())

	return nil
}