package main

// generate style helper from taken from
// https://www.w3.org/Style/CSS/all-properties

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strings"
	"text/template"
)

var fileTmpl = template.Must(template.New("file").Parse(`
// +build js,wasm

package wasm

/*
https://www.w3.org/Style/CSS/all-properties#list
*/

type cssStyleHelper interface {
	{{ .Methods }}
}

type cssStyleHelperImpl struct {
	*cssStyleDeclarationImpl
}

func newCSSStyleHelperImpl(v *cssStyleDeclarationImpl) *cssStyleHelperImpl {
	if v.Valid() {
		return &cssStyleHelperImpl {
			cssStyleDeclarationImpl: v,
		}
	}
	return nil
}

{{.Implementation}}
`))

var methodTmpl = template.Must(template.New("method").Parse(`
	{{ .UpName }}() string
	Set{{.UpName -}}(string)`))

var implTmpl = template.Must(template.New("style").Parse(`
func (p *cssStyleHelperImpl) {{.UpName}}() string {
	return p.PropertyValue("{{.Name}}")
}

func (p *cssStyleHelperImpl) Set{{.UpName}}(s string) {
	p.SetProperty("{{.Name}}",s)
}
`))

type style struct {
	Name   string
	UpName string
}

var styles []style

func main() {
	f, err := os.Open("styles.txt")
	if err != nil {
		log.Fatal("styles.txt file not found")
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		s := style{
			Name:   line,
			UpName: toUpper(line),
		}
		styles = append(styles, s)
	}

	var (
		implBuffer   bytes.Buffer
		methodBuffer bytes.Buffer
	)

	for _, s := range styles {
		implTmpl.Execute(&implBuffer, s)
		methodTmpl.Execute(&methodBuffer, s)
	}

	fileTmpl.Execute(os.Stdout, struct {
		Methods        string
		Implementation string
	}{
		Methods:        methodBuffer.String(),
		Implementation: implBuffer.String(),
	})
}

func toUpper(s string) string {
	strs := strings.Split(s, "-")
	for i, str := range strs {
		strs[i] = strings.Title(str)
	}
	return strings.Join(strs, "")
}
