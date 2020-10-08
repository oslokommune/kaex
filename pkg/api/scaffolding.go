package api

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"
)

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

	if res.StatusCode != 200 {
		return "", fmt.Errorf("%s not found", url)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return "", readErr
	}

	return string(body), nil
}

func fetchLocalTemplate(kaex Kaex, writer io.Writer, name string) error {
	templatesPath := path.Join(kaex.ConfigPath, "templates")

	mainPath := path.Join(templatesPath, name)
	alternativePath := path.Join(templatesPath, fmt.Sprintf("%s.yaml", name))

	var targetPath string
	_, mainPathErr := os.Stat(mainPath)
	_, alternativePathErr := os.Stat(alternativePath)

	if mainPathErr != nil && alternativePathErr != nil {
		return fmt.Errorf("could not find %s/%s locally: %w", mainPath, alternativePath, mainPathErr)
	}

	switch {
	case mainPathErr == nil:
		targetPath = mainPath
	case alternativePathErr == nil:
		targetPath = alternativePath
	}

	targetFile, err := os.Open(targetPath)
	if err != nil {
		err = fmt.Errorf("error while reading %s: %w", targetPath, err)

		return err
	}

	_, err = io.Copy(writer, targetFile)
	if err != nil {
		return fmt.Errorf("error writing contents of %s to writer: %w", targetPath, err)
	}

	return nil
}

func fetchRemoteTemplate(kaex Kaex, w io.Writer, name string) error {
	rawRepoURL, err := url.Parse(kaex.RepoURL)
	if err != nil {
		return fmt.Errorf("malformed url %s: %w", kaex.RepoURL, err)
	}

	rawRepoURL.Path = path.Join(rawRepoURL.Path, "examples")

	mainURL, _ := url.Parse(rawRepoURL.String())
	alternativeURL, _ := url.Parse(rawRepoURL.String())

	mainURL.Path = path.Join(mainURL.Path, name)
	alternativeURL.Path = path.Join(alternativeURL.Path, fmt.Sprintf("%s.yaml", name))

	mainTemplate, mainErr := fetchURL(mainURL.String())
	alternativeTemplate, alternativeErr := fetchURL(alternativeURL.String())
	if mainErr != nil && alternativeErr != nil {
		return fmt.Errorf("unable to fetch remote template: %w", mainErr)
	}

	switch {
	case mainErr == nil:
		fmt.Fprint(w, mainTemplate)
	case alternativeErr == nil:
		fmt.Fprint(w, alternativeTemplate)
	}

	return nil
}

func FetchTemplate(kaex Kaex, writer io.Writer, name string) error {
	if kaex.ConfigPath != "" {
		err := fetchLocalTemplate(kaex, writer, name)
		if err == nil {
			return nil
		}
	}

	err := fetchRemoteTemplate(kaex, writer, name)
	if err == nil {
		return nil
	}

	return fmt.Errorf("unable to find template %s locally nor remotely: %w", name, err)
}
