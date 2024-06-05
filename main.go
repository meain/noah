package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"
	"time"

	"golang.org/x/net/html"
)

const (
	dateFormat = "2006-01-02"
)

func printError(msg string) {
	os.Stderr.WriteString("[ERROR]" + msg)
	os.Exit(2)
}

func main() {
	if len(os.Args) != 2 {
		os.Stderr.WriteString("Usage: noah <url>")
		os.Exit(1)
	}

	url := os.Args[1]
	templateType := getTempplateType(url)
	templatePath := getTemplate(templateType)
	templateName := path.Base(templatePath)

	tmpl, err := template.New(templateName).ParseFiles(templatePath)
	if err != nil {
		printError(err.Error())
	}

	data, err := getData(url, templateType)
	if err != nil {
		printError(err.Error())
	}

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		printError(err.Error())
	}
}

func getTempplateType(_ string) string {
	// TODO(meain): parse from url
	return TemplateTypeArticle
}

func getTemplate(templateType string) string {
	// TODO(meain): read from config
	var tmplFile = "templates/" + templateType + ".tmpl"
	return tmplFile
}

func getData(url string, templateType string) (map[string]string, error) {
	switch templateType {
	case TemplateTypeArticle:
		return getArticleData(url)
	}
	return nil, fmt.Errorf("unknown template type: %s", templateType)
}

func getArticleData(url string) (map[string]string, error) {
	// fetch html from page
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	title := getTitle(doc)
	return map[string]string{
		"Title": title,
		"URL":   url,
		"Date":  time.Now().Format(dateFormat),
	}, nil
}

func getTitle(htmlNode *html.Node) string {
	if htmlNode.Type == html.ElementNode && htmlNode.Data == "title" {
		if htmlNode.FirstChild != nil {
			return htmlNode.FirstChild.Data
		}
	}
	for child := htmlNode.FirstChild; child != nil; child = child.NextSibling {
		title := getTitle(child)
		if len(title) != 0 {
			return title
		}
	}
	return ""
}
