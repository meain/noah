package main

// TODO(meain): make content converted to markdown available (use ff readbility)
// Also have to have a documentation of the supported fields somewhere

import (
	"bytes"
	"io"
	"net/http"
	"net/url"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"golang.org/x/net/html"
)

func getArticleData(input string) (map[string]string, error) {
	parsedURL, err := url.Parse(input)
	if err != nil {
		return nil, err
	}

	host := parsedURL.Host

	resp, err := http.Get(input)
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

	return map[string]string{
		"Title":           getTitle(doc),
		"URL":             input,
		"Host":            host,
		"Author":          getAuthor(doc),
		"HTMLContent":     string(content),
		"MarkdownContent": markdown,
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
