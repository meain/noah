package main

const (
	TemplateTypeArticle = "article"
	TemplateTypeYoutube = "youtube"
)

type Item interface {
	// Get the data associated with the thing
	getData() (map[string]string, error)

	// Pas in the data to get a filename to save
	getFileName(map[string]string) string
}
