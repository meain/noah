package main

const (
	TemplateTypeArticle = "article"
	TemplateTypeYoutube = "youtube"
)

type Item interface {
	// Get the data associated with the thing
	getData() (map[string]any, error)

	// Pas in the data to get a filename to save
	getFileName(map[string]any) string
}

type YTDlpResponse struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Thumbnails []struct {
		URL        string `json:"url"`
		Preference int    `json:"preference"`
		ID         string `json:"id"`
		Height     int    `json:"height,omitempty"`
		Width      int    `json:"width,omitempty"`
		Resolution string `json:"resolution,omitempty"`
	} `json:"thumbnails"`
	Thumbnail        string   `json:"thumbnail"`
	Description      string   `json:"description"`
	ChannelID        string   `json:"channel_id"`
	ChannelURL       string   `json:"channel_url"`
	Duration         int      `json:"duration"`
	ViewCount        int      `json:"view_count"`
	AverageRating    any      `json:"average_rating"`
	AgeLimit         int      `json:"age_limit"`
	WebpageURL       string   `json:"webpage_url"`
	Categories       []string `json:"categories"`
	Tags             []string `json:"tags"`
	PlayableInEmbed  bool     `json:"playable_in_embed"`
	LiveStatus       string   `json:"live_status"`
	ReleaseTimestamp any      `json:"release_timestamp"`
	FormatSortFields []string `json:"_format_sort_fields"`
	Subtitles        struct {
	} `json:"subtitles"`
	CommentCount int `json:"comment_count"`
	Chapters     []struct {
		StartTime float64 `json:"start_time"`
		Title     string  `json:"title"`
		EndTime   float64 `json:"end_time"`
	} `json:"chapters"`
	Heatmap              any     `json:"heatmap"`
	LikeCount            int     `json:"like_count"`
	Channel              string  `json:"channel"`
	ChannelFollowerCount int     `json:"channel_follower_count"`
	ChannelIsVerified    bool    `json:"channel_is_verified"`
	Uploader             string  `json:"uploader"`
	UploaderID           string  `json:"uploader_id"`
	UploaderURL          string  `json:"uploader_url"`
	UploadDate           string  `json:"upload_date"`
	Timestamp            int     `json:"timestamp"`
	Availability         string  `json:"availability"`
	OriginalURL          string  `json:"original_url"`
	WebpageURLBasename   string  `json:"webpage_url_basename"`
	WebpageURLDomain     string  `json:"webpage_url_domain"`
	Extractor            string  `json:"extractor"`
	ExtractorKey         string  `json:"extractor_key"`
	Playlist             any     `json:"playlist"`
	PlaylistIndex        any     `json:"playlist_index"`
	DisplayID            string  `json:"display_id"`
	Fulltitle            string  `json:"fulltitle"`
	DurationString       string  `json:"duration_string"`
	ReleaseYear          any     `json:"release_year"`
	IsLive               bool    `json:"is_live"`
	WasLive              bool    `json:"was_live"`
	RequestedSubtitles   any     `json:"requested_subtitles"`
	HasDrm               any     `json:"_has_drm"`
	Epoch                int     `json:"epoch"`
	Format               string  `json:"format"`
	FormatID             string  `json:"format_id"`
	Ext                  string  `json:"ext"`
	Protocol             string  `json:"protocol"`
	Language             string  `json:"language"`
	FormatNote           string  `json:"format_note"`
	FilesizeApprox       int     `json:"filesize_approx"`
	Tbr                  float64 `json:"tbr"`
	Width                int     `json:"width"`
	Height               int     `json:"height"`
	Resolution           string  `json:"resolution"`
	Fps                  int     `json:"fps"`
	DynamicRange         string  `json:"dynamic_range"`
	Vcodec               string  `json:"vcodec"`
	Vbr                  float64 `json:"vbr"`
	StretchedRatio       any     `json:"stretched_ratio"`
	AspectRatio          float64 `json:"aspect_ratio"`
	Acodec               string  `json:"acodec"`
	Abr                  float64 `json:"abr"`
	Asr                  int     `json:"asr"`
	AudioChannels        int     `json:"audio_channels"`
	Filename             string  `json:"_filename"`
	Filename0            string  `json:"filename"`
	Type                 string  `json:"_type"`
	Version              struct {
		Version        string `json:"version"`
		CurrentGitHead any    `json:"current_git_head"`
		ReleaseGitHead string `json:"release_git_head"`
		Repository     string `json:"repository"`
	} `json:"_version"`
}
