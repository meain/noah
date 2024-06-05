package main

import (
	"fmt"
	"html/template"
	"os"
	"path"
	"time"
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

	// Augment initial data
	data["Date"] = time.Now().Format(dateFormat)

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
