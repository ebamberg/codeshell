package templating

import (
	"bytes"
	"codeshell/output"
	"html/template"
)

func ProcessPlaceholders(withPlaceholders string, context any) string {
	varTemplate, err := template.New("new").Parse(withPlaceholders)
	if err == nil {
		buf := bytes.NewBufferString("")
		varTemplate.Execute(buf, context)
		return buf.String()
	} else {
		output.Errorln(err)
		return withPlaceholders
	}

}
