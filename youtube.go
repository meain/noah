package main

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"time"
)

type NoEmbedResponse struct {
	AuthorName      string `json:"author_name"`
	AuthorURL       string `json:"author_url"`
	ThumbnailURL    string `json:"thumbnail_url"`
	Title           string `json:"title"`
	ProviderURL     string `json:"provider_url"`
	Height          int    `json:"height"`
	Type            string `json:"type"`
	HTML            string `json:"html"`
	Width           int    `json:"width"`
	ThumbnailWidth  int    `json:"thumbnail_width"`
	ProviderName    string `json:"provider_name"`
	URL             string `json:"url"`
	ThumbnailHeight int    `json:"thumbnail_height"`
	Version         string `json:"version"`
}

const noEmbedUrl = "https://noembed.com/embed?dataType=json&url="

type youTubeItem struct {
	URL string
}

func (y youTubeItem) getData() (map[string]string, error) {
	data := map[string]string{"URL": y.URL}

	resp, err := http.Get(noEmbedUrl + y.URL)
	if err != nil {
		return data, err
	}

	defer resp.Body.Close()

	noer := NoEmbedResponse{}
	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&noer)
	if err != nil {
		return data, err
	}

	data["Channel"] = noer.AuthorName
	data["ChannelURL"] = noer.AuthorURL
	data["Title"] = noer.Title
	data["Thumbnail"] = noer.ThumbnailURL

	return data, nil
}

func (y youTubeItem) getFileName(data map[string]string) string {
	folder := "Youtube"
	fileName := time.Now().Format("2006-01-02 15:04:05")

	title, ok := data["Title"]
	if ok {
		fileName = title
	}

	return filepath.Join(folder, fileName+".md")
}
