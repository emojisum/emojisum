// +build ignore

package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"text/template"

	"github.com/vbatts/emojisum/emoji"
)

func main() {
	flag.Parse()
	input, err := os.Open(*flInput)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	vm := emoji.VersionedMap{}

	dec := json.NewDecoder(input)
	if err := dec.Decode(&vm); err != nil {
		log.Fatal(err)
	}

	output, err := os.Create(*flOutput)
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	mapGoTemp := template.Must(template.ParseFiles(*flTemplate))
	if err := mapGoTemp.Execute(output, vm); err != nil {
		log.Fatal(err)
	}
}

var (
	flInput    = flag.String("in", "emojimap.json", "json input")
	flOutput   = flag.String("out", "map_gen.go", "golang output")
	flTemplate = flag.String("template", "map_gen.tmpl", "template of golang source to use")
)
