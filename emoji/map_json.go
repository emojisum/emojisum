// +build ignore

package main

import (
	"encoding/json"
	"log"
	"os"
	"text/template"
)

func main() {
	input, err := os.Open("map-draft.json")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	// these are an ordered list, referened by a byte (each byte of a checksum digest)
	Map := []string{}

	dec := json.NewDecoder(input)
	if err := dec.Decode(&Map); err != nil {
		log.Fatal(err)
	}

	output, err := os.Create("map_gen.go")
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()
	if err := mapGoTemp.Execute(output, Map); err != nil {
		log.Fatal(err)
	}
}

var (
	mapGoText = `// THIS FILE IS GENERATED. DO NOT EDIT.

package emoji

var sumList = []string{ {{- range . }}
	"{{.}}",{{- end }}
}
`
	mapGoTemp = template.Must(template.New("map.go").Parse(mapGoText))
)
