package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"text/template"
)

type TemplateData struct {
	PackageName string
	TestName    string
}

func writeTemplate(w io.Writer, templateText string, data TemplateData) error {

	tmpl, err := template.New("go-file").Parse(templateText)
	if err != nil {
		return fmt.Errorf("Template konnte nicht verarbeitet werden: %w", err)
	}

	var buffer bytes.Buffer

	if err := tmpl.Execute(&buffer, data); err != nil {
		return fmt.Errorf("Template konnte nicht generiert werden: %w", err)
	}

	formattedCode, err := format.Source(buffer.Bytes())
	if err != nil {
		return fmt.Errorf("Go-Code konnte nicht formatiert werden: %w", err)
	}

	if _, err := w.Write(formattedCode); err != nil {
		return fmt.Errorf("Formatierter Code konnte nicht geschrieben werden: %w", err)
	}

	return nil
}

const goFileTemplate = `package {{.PackageName}}
`

const goTestTemplate = `package {{.PackageName}}

import "testing"

func Test{{.TestName}}(t *testing.T) {
	// TODO: Test implementieren
}
`
