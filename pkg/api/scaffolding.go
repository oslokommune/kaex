package api

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func FetchMinimalExample(w io.Writer) error {
	url := "https://raw.githubusercontent.com/deifyed/kaex/master/examples/application-minimal.yaml"
	
	example, err := fetchExample(url)
	if err != nil {
		return err
	} 
	
	fmt.Fprintf(w, example)
	
	return nil
}

func FetchFullExample(w io.Writer) error {
	url := "https://raw.githubusercontent.com/deifyed/kaex/master/examples/application-full.yaml"

	example, err := fetchExample(url)
	if err != nil {
		return err
	}

	fmt.Fprintf(w, example)

	return nil
}

func fetchExample(url string) (string, error) {
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
