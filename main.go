package main

import (
	"fmt"
	"os"
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

	tmpl, err := getTemplate(templateType)
	if err != nil {
		printError(err.Error())
	}

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		printError(err.Error())
	}
}

func getData(url string, templateType string) (map[string]string, error) {
	switch templateType {
	case TemplateTypeArticle:
		return getArticleData(url)
	}
	return nil, fmt.Errorf("unknown template type: %s", templateType)
}
