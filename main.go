package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/alecthomas/kong"
)

const (
	dateFormat = "2006-01-02"
)

func printError(msg string) {
	os.Stderr.WriteString("[ERROR]" + msg)
	os.Exit(2)
}

var cli struct {
	Input  string `arg:"" optional:"" name:"input" help:"Item to fetch metadata"`
	Output string `optional:"" short:"o" name:"output" help:"Directory to write out metadata file"`
}

func main() {
	ctx := kong.Parse(&cli)
	switch ctx.Command() {
	case "<input>":
		doIt(cli.Input, cli.Output)
	default:
		os.Stderr.WriteString("Usage: noah [arguments] <input>")
		os.Exit(1)
	}
}

func doIt(input, outDir string) {
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

	outFile := os.Stdout
	if len(outDir) != 0 {
		err := os.MkdirAll(outDir, os.ModePerm)
		if err != nil {
			printError(err.Error())
		}

		outFile, err = os.Create(filepath.Join(outDir, getFileName(templateType, data)))
		if err != nil {
			printError(err.Error())
		}

		os.Stdout.WriteString(fmt.Sprintf("Writing to %s\n", outFile.Name()))
		defer outFile.Close()
	}

	err = tmpl.Execute(outFile, data)
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
