package main

import (
	_ "embed"
	"net/url"
	"text/template"
)

//go:embed templates/article.tmpl
var articleTemplate string

//go:embed templates/youtube.tmpl
var youtubeTemplate string

func getTemplateType(input string) string {
	purl, err := url.Parse(input)
	if err != nil {
		return TemplateTypeArticle
	}

	if purl.Host == "www.youtube.com" ||
		purl.Host == "youtube.com" ||
		purl.Host == "youtu.be" {
		return TemplateTypeYoutube
	}

	return TemplateTypeArticle
}

func getTemplate(templateType string) (*template.Template, error) {
	// TODO(meain): read from config if specified
	switch templateType {
	case TemplateTypeYoutube:
		return template.New(templateType).Parse(youtubeTemplate)
	default:
		return template.New(templateType).Parse(articleTemplate)
	}
}
