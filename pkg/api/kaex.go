package api

import "io"

type Kaex struct {
	Err io.Writer
	Out io.Writer
	In io.Reader

	ConfigPath string
	TemplatesDirURL string
}
