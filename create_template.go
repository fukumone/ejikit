package ejikit

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"log"
	"os"
	"text/template"
)

type GemojiEmoji struct {
	Aliases     []string `json:"aliases"`
	Description string   `json:"description"`
	Emoji       string   `json:"emoji"`
	Tags        []string `json:"tags"`
}

type TemplateData struct {
	PkgName string
	CodeMap map[string]string
}

const templateMapCode = `
package {{.PkgName}}

// Mapping from character to concrete escape code.
var emojiCodeMap = map[string]string{
  {{range $key, $val := .CodeMap}}":{{$key}}:": {{$val}},
{{end}}
}

var emojiNameArray = []string{
  {{range $key, $val := .CodeMap}}":{{$key}}:",
{{end}}
}

var emojiValueArray = []string{
	{{range $key, $val := .CodeMap}}{{$val}},
{{end}}
}
`

var pkgName string
var fileName string

func init() {
	log.SetFlags(log.Llongfile)

	flag.StringVar(&pkgName, "pkg", "main", "output package")
	flag.StringVar(&fileName, "o", "emoji_template.go", "output file")
	flag.Parse()
}

func CreateTemplate() {

	emojiMap, err := GenerateJson(true)
	if err != nil {
		log.Fatalln(err)
	}

	var buf bytes.Buffer
	t := template.Must(template.New("template").Parse(templateMapCode))
	if err := t.Execute(&buf, TemplateData{pkgName, emojiMap}); err != nil {
		log.Fatalln(err)
	}

	bts, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Println(string(buf.Bytes()))
		log.Fatalln(err)
	}

	os.Remove(fileName)

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	if _, err := file.Write(bts); err != nil {
		log.Fatalln(err)
	}
}
