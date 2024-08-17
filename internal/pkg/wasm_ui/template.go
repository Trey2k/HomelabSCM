package wasm_ui

import (
	"strings"
	"text/template"
)

func ParseTemplate(src, templateName string, templateData any) (string, error) {
	var htmlBuilder strings.Builder
	tpl, err := template.New("").Parse(src)
	if err != nil {
		return "", err
	}
	
	err = tpl.ExecuteTemplate(&htmlBuilder, templateName, templateData)
	if err != nil {
		return "", err
	}

	return htmlBuilder.String(), nil
}