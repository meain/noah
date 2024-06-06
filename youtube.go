package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
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

func (y youTubeItem) getData() (map[string]any, error) {
	data, err := getDataFromYtDlp(y.URL)
	if err != nil {
		fmt.Println("youtube.go:37 err:", err)
		return getDataFromNoEmbed(y.URL)
	}

	return data, nil
}

func getDataFromYtDlp(input string) (map[string]any, error) {
	// check if yt-dlp available in PATH
	_, err := exec.LookPath("yt-dlp")
	if err != nil {
		return map[string]any{}, err
	}

	cmd := exec.Command("yt-dlp", "--print-json", "--skip-download", input)
	out, err := cmd.Output()
	if err != nil {
		return map[string]any{}, err
	}

	data := map[string]any{"URL": input}

	var resp YTDlpResponse

	err = json.Unmarshal(out, &resp)
	if err != nil {
		return data, err
	}

	data["Channel"] = resp.Channel
	data["ChannelURL"] = resp.ChannelURL
	data["Title"] = resp.Title
	data["Thumbnail"] = resp.Thumbnail
	data["Chapters"] = resp.Chapters
	data["Views"] = resp.ViewCount
	data["Likes"] = resp.LikeCount
	data["Description"] = resp.Description
	data["Duration"] = resp.Duration
	data["Categories"] = resp.Categories
	data["CommentCount"] = resp.CommentCount
	// TODO: add more as required

	return data, nil
}

func getDataFromNoEmbed(input string) (map[string]any, error) {
	data := map[string]any{"URL": input}

	resp, err := http.Get(noEmbedUrl + input)
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

func (y youTubeItem) getFileName(data map[string]any) string {
	folder := "Youtube"
	fileName := time.Now().Format("2006-01-02 15:04:05")

	title, ok := data["Title"].(string)
	if ok {
		fileName = title
	}

	return filepath.Join(folder, fileName+".md")
}
