package main

import (
	_ "embed"
	"text/template"
)

//go:embed templates/article.tmpl
var articleTemplate string

func getTempplateType(_ string) string {
	// TODO(meain): parse from url
	return TemplateTypeArticle
}

func getTemplate(templateType string) (*template.Template, error) {
	// TODO(meain): read from config if specified
	return template.New(templateType).Parse(articleTemplate)
}
