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
		os.Stderr.WriteString("Usage: noah <input>")
		os.Exit(1)
	}

	input := os.Args[1]
	templateType := getTempplateType(input)
	templatePath := getTemplate(templateType)
	templateName := path.Base(templatePath)

	tmpl, err := template.New(templateName).ParseFiles(templatePath)
	if err != nil {
		printError(err.Error())
	}

	data, err := getData(input, templateType)
	if err != nil {
		printError(err.Error())
	}

	// Augment initial data
	data["Input"] = input
	data["Date"] = time.Now().Format(dateFormat)
	data["Time"] = time.Now().Format(time.Kitchen)
	data["ISODateTime"] = time.Now().Format(time.RFC3339)
	data["TemplateType"] = templateType

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
