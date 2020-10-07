package api

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"time"
)

const (
	REPO_EXAMPLE_BASE_URL = "https://raw.githubusercontent.com/oslokommune/kaex/master/examples"
)

func FetchMinimalExample(w io.Writer) error {
	exampleFileName := "application-minimal.yaml"

	err := fetchExample(w, exampleFileName)
	if err != nil {
		err = fmt.Errorf("unable to fetch %s: %w", exampleFileName, err)

		return err
	}

	return nil
}

func FetchFullExample(w io.Writer) error {
	exampleFileName := "application-full.yaml"

	err := fetchExample(w, exampleFileName)
	if err != nil {
		err = fmt.Errorf("unable to fetch %s: %w", exampleFileName, err)

		return err
	}

	return nil
}

func fetchLocalExample(name string) (string, error) {
	configPath, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	targetPath := path.Join(configPath, fmt.Sprintf("kaex/%s", name))
	_, err = os.Stat(targetPath)
	if err != nil {
		err = fmt.Errorf("no such file/path \"%s\": %w", targetPath, err)

		return "", err
	}

	data, err := ioutil.ReadFile(targetPath)
	if err != nil {
		err = fmt.Errorf("error while reading file: %w", err)

		return "", err
	}

	return string(data), nil
}

func fetchExample(w io.Writer, name string) error {
	example, err := fetchLocalExample(name)
	if err == nil {
		fmt.Fprintf(w, example)

		return nil
	}

	url := fmt.Sprintf("%s/%s", REPO_EXAMPLE_BASE_URL, name)

	example, err = fetchURL(url)
	if err != nil {
		return err
	}

	fmt.Fprintf(w, example)

	return nil
}

func fetchURL(url string) (string, error) {
	spaceClient := http.Client{
		Timeout: 2 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		return "", getErr
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return "", readErr
	}

	return string(body), nil
}

func GenerateDebugger(writer io.Writer) error {
	app := Application{
		Name:            "debugger",
		Image:           "markeijsermans/debug",
		Version:         "kitchen-sink",
		Replicas:        1,
	}

	pod, err := CreatePod(app)
	if err != nil {
		return fmt.Errorf("error creating debugger pod: %w", err)
	}

	pod.Spec.Containers[0].Command = []string{"/bin/sh", "-c", "sleep 3600"}

	err = WriteResource(writer, pod)
	if err != nil {
		return fmt.Errorf("error writing resource to writer: %w", err)
	}

	return nil
}
