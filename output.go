package main

import (
	"path/filepath"
	"time"
)

// TODO(meain): this should be configurable
func getFileName(templateType string, data map[string]string) string {
	folder := ""
	fileName := time.Now().Format("2006-01-02 15:04:05")
	switch templateType {
	case TemplateTypeArticle:
		folder = "Article"
		title, ok := data["Title"]
		if ok {
			fileName = title
		}
	}

	return filepath.Join(folder, fileName+".md")
}
