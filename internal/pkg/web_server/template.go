package web_server

import (
	"fmt"
	"html/template"
	"net/http"

	web_fs "homelabscm.com/scm/web"
)

func executeTemplate(name string, w http.ResponseWriter, data any) error {
	t, err := template.New("").ParseFS(web_fs.TemplateFS,
		fmt.Sprintf("template/%s.html", name),
		"template/base.html")
	if err != nil {
		return err
	}

	return t.ExecuteTemplate(w, "base", data)
}
