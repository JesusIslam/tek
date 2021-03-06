// +build ignore

// This program generates indonesian_pos.go
// It can be invoked by running go generate

package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"text/template"
)

const FileTemplate = `// This file is generated by go generate, please do not change manually.

package tek

var {{ .VariableName }} = []*Vocab{
	{{ range .Data }}{
		Id: {{ .Id }},
		Word: "{{ .Word }}",
		Type: "{{ .Type }}",
	},{{ end }}
}
`

type TemplateFiller struct {
	VariableName string
	Data         []*Vocab
}

type Vocab struct {
	Id   int    `json:"id"`
	Word string `json:"word"`
	Type string `json:"type"`
}

func main() {
	pos := []*Vocab{}
	fb, err := ioutil.ReadFile("./pos_id.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(fb, &pos)
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("indonesian_pos.go").Parse(FileTemplate)
	if err != nil {
		panic(err)
	}

	targetFile, err := os.Create("./indonesian_pos.go")
	if err != nil {
		panic(err)
	}
	defer targetFile.Close()

	templateFiller := &TemplateFiller{
		VariableName: "indonesianPos",
		Data:         pos,
	}

	err = tmpl.Execute(targetFile, templateFiller)
	if err != nil {
		panic(err)
	}
}
