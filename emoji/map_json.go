// +build ignore

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/emojisum/emojisum/emoji"
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

	//mapGoTemp := template.Must(template.ParseFiles(*flTemplate))
	funcMap := template.FuncMap{
		"IsColonNotation":       emoji.IsColonNotation,
		"IsCodepoint":           emoji.IsCodepoint,
		"CodepointLinkMarkdown": codepointLinkMarkdown,
	}

	mapGoTemp, err := template.New("").Funcs(funcMap).Parse(tmpl[*flTemplate])
	if err != nil {
		log.Fatal(err)
	}
	if err := mapGoTemp.Execute(output, vm); err != nil {
		log.Fatal(err)
	}
}

func codepointLinkMarkdown(word string) string {
	return fmt.Sprintf(`[%s](%s)`, word, emoji.UnicodeLinkURL(word))
}

var (
	flInput    = flag.String("in", "emojimap.json", "json input")
	flOutput   = flag.String("out", "map_gen.go", "golang output")
	flTemplate = flag.String("template", "map_gen", "template to use (map_gen or markdown_gen)")
)

var tmpl = map[string]string{
	"map_gen": `
// THIS FILE IS GENERATED. DO NOT EDIT.

package emoji

func init() {
  mapGen = VersionedMap{
	  Description: "{{.Description}}",
	  Version: "{{.Version}}",
	  EmojiWords: []Words{ {{- range .EmojiWords }}
      Words{ {{ range . -}}
        "{{- . }}",{{- end }}
      },{{- end }}
	  },
  }
}
`,
	"markdown_gen": `
## Emoji Map list

_THIS FILE IS GENERATED. DO NOT EDIT._

This is for "pretty" viewing purposes.
To view the functional document, see [emojimap.json](./emojimap.json).

### Description

{{ .Description }}

### Version

{{ .Version }}

### List

{{- range $index, $words := .EmojiWords }}
  * ` + "`{{ $index }}`" + ` -- {{ range $words }} {{- if IsColonNotation . -}} {{ . }} ` + "`{{ . }}`" + ` {{- else }} {{ CodepointLinkMarkdown . }} {{- end }}{{- end }}
{{- end }}
`,
}
