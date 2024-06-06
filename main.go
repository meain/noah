package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/alecthomas/kong"
)

const (
	dateFormat = "2006-01-02"
)

func printError(msg string) {
	os.Stderr.WriteString("[ERROR] " + msg + "\n")
	os.Exit(2)
}

var cli struct {
	Input  string `arg:"" optional:"" name:"input" help:"Item to fetch metadata"`
	Force  bool   `optional:"" short:"f" name:"force" help:"Overwrite existing files"`
	Output string `optional:"" short:"o" name:"output" help:"Directory to write out metadata file"`
}

func main() {
	ctx := kong.Parse(&cli)
	switch ctx.Command() {
	case "<input>":
		doIt(cli.Input, cli.Output, cli.Force)
	default:
		os.Stderr.WriteString("Usage: noah [arguments] <input>")
		os.Exit(1)
	}
}

func doIt(input, outDir string, force bool) {
	templateType := getTemplateType(input)

	it, err := getItem(input, templateType)
	if err != nil {
		printError(err.Error())
	}

	data, err := it.getData()
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
		// fileName might also contain a dir
		fileName := filepath.Clean(it.getFileName(data))

		// Cleanup the filename
		invalidChars := regexp.MustCompile(`[^a-zA-Z0-9_\-\. :]`)
		fileName = invalidChars.ReplaceAllString(fileName, "_")

		fullPath := filepath.Join(outDir, fileName)

		if _, err := os.Stat(fullPath); err == nil {
			if !force {
				os.Stdout.WriteString(fmt.Sprintf("File already exists: %s\n", filepath.Base(fullPath)))
				os.Exit(0)
			}
		}

		err := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm)
		if err != nil {
			printError(err.Error())
		}

		outFile, err = os.Create(fullPath)
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

func getItem(url string, templateType string) (Item, error) {
	switch templateType {
	case TemplateTypeArticle:
		return articleItem{url}, nil
	case TemplateTypeYoutube:
		return youTubeItem{url}, nil
	}
	return nil, fmt.Errorf("unknown template type: %s", templateType)
}
