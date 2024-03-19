package main

import (
	"embed"
	"fmt"
	"mtechnavi/sharelib/cmd/protoc-gen-csv/funcs"
	"text/template"

	"google.golang.org/protobuf/compiler/protogen"
)

//go:embed assets/*.tmpl
var assets embed.FS

func generateFile(p *protogen.Plugin, file *protogen.File) {
	outName := file.GeneratedFilenamePrefix + ".pb.csv.go"
	tpl, err := template.New(outName).
		Funcs(funcs.Map).
		ParseFS(assets, "assets/*.tmpl")
	if err != nil {
		p.Error(err)
		return
	}

	// CSV用の指定があるもののみ、ファイル生成の対象とする
	var messages []*protogen.Message
	for _, m := range file.Messages {
		for _, f := range m.Fields {
			_, ok := funcs.GetFieldExtension(f.Desc)
			if !ok {
				continue
			}
			messages = append(messages, m)
			break
		}
	}
	if len(messages) == 0 {
		return
	}

	// validate
	valid := true
	for _, m := range messages {
		for _, errMsg := range funcs.ValidateMessage(m) {
			p.Error(fmt.Errorf("%s: %s", m.GoIdent.GoName, errMsg))
		}
	}
	if !valid {
		return
	}

	w := p.NewGeneratedFile(outName, file.GoImportPath)
	if err := tpl.ExecuteTemplate(w, "generated-file.go.tmpl", map[string]any{
		"Plugin":   p,
		"File":     file,
		"Messages": messages,
	}); err != nil {
		p.Error(err)
		return
	}
}

func main() {
	var opts protogen.Options
	opts.Run(func(p *protogen.Plugin) error {
		for _, f := range p.Files {
			if !f.Generate {
				continue
			}
			generateFile(p, f)
		}
		return nil
	})
}
