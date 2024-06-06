package main

// TODO(meain): make content converted to markdown available (use ff readbility)
// Also have to have a documentation of the supported fields somewhere

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"time"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"golang.org/x/net/html"
)

type articleItem struct {
	URL string
}

func (a articleItem) getData() (map[string]any, error) {
	parsedURL, err := url.Parse(a.URL)
	if err != nil {
		return nil, err
	}

	host := parsedURL.Host

	resp, err := http.Get(a.URL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	doc, err := html.Parse(bytes.NewReader(content))
	if err != nil {
		return nil, err
	}

	converter := md.NewConverter("", true, nil)

	markdown, err := converter.ConvertString(string(content))
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"Title":           getTitle(doc),
		"URL":             a.URL,
		"Host":            host,
		"Author":          getAuthor(doc),
		"HTMLContent":     string(content),
		"MarkdownContent": markdown,
	}, nil
}

func (y articleItem) getFileName(data map[string]any) string {
	folder := "Article"
	fileName := time.Now().Format("2006-01-02 15:04:05")

	title, ok := data["Title"]
	if ok {
		fileName = title.(string)
	}

	return filepath.Join(folder, fileName+".md")
}

func getTitle(htmlNode *html.Node) string {
	if htmlNode.Type == html.ElementNode && htmlNode.Data == "title" {
		if htmlNode.FirstChild != nil {
			return maybeURLDecode(htmlNode.FirstChild.Data)
		}
	}
	for child := htmlNode.FirstChild; child != nil; child = child.NextSibling {
		title := getTitle(child)
		if len(title) != 0 {
			return maybeURLDecode(title)
		}
	}
	return ""
}

func getAuthor(htmlNode *html.Node) string {
	if htmlNode.Type == html.ElementNode && htmlNode.Data == "meta" {
		found := false
		content := ""
		for _, attr := range htmlNode.Attr {
			if attr.Key == "name" && attr.Val == "author" {
				found = true
			}
			if attr.Key == "content" {
				content = attr.Val
			}
		}

		if found {
			return content
		}
	}
	for child := htmlNode.FirstChild; child != nil; child = child.NextSibling {
		author := getAuthor(child)
		if len(author) != 0 {
			return author
		}
	}
	return ""
}
