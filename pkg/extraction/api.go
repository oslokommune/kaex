package extraction

import (
	"errors"
	"regexp"
	"strings"
)

var ErrorNoMatch = errors.New("no match found for kind")

var kindRetriever = regexp.MustCompile(`kind:\s(\w+)`)

type Resource struct {
	Kind string
	Data []byte
}

func ExtractKind(data []byte) (string, error) {
	match := kindRetriever.FindSubmatch(data)

	if match == nil {
		return "", ErrorNoMatch
	}

	return strings.ToLower(string(match[1])), nil
}
