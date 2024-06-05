package main

import (
	"net/http"
	"time"

	"golang.org/x/net/html"
)

func getArticleData(url string) (map[string]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"Title":  getTitle(doc),
		"URL":    url,
		"Date":   time.Now().Format(dateFormat),
		"Author": getAuthor(doc),
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
