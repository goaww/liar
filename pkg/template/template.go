package template

import (
	"bytes"
	"github.com/google/uuid"
	"html/template"
)

type Template interface {
	RenderFromFile(path string, ctx Context) (error, string)

	Render(text string, ctx Context) (error, string)
}

type HtmlTemplate struct {
}

func NewHtmlTemplate() *HtmlTemplate {
	return &HtmlTemplate{}
}

func (h *HtmlTemplate) RenderFromFile(path string, ctx Context) (error, string) {
	t, err := template.ParseFiles(path)
	if err != nil {
		return err, ""
	} else {

		var tpl bytes.Buffer
		if err := t.Execute(&tpl, ctx); err != nil {
			return err, ""
		} else {
			return nil, tpl.String()
		}

	}
}

func (h *HtmlTemplate) Render(text string, ctx Context) (error, string) {
	t, err := template.New(uuid.New().String()).Parse(text)
	if err != nil {
		return err, ""
	} else {

		var tpl bytes.Buffer
		if err := t.Execute(&tpl, ctx); err != nil {
			return err, ""
		} else {
			return nil, tpl.String()
		}

	}
}
