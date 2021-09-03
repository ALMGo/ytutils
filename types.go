package ytutils

type SearchOptions struct {
	PageToken  string
	Order      string
	MaxResults int64
	SearchType []string
}

type Thumbnails struct {
	MaxRes      string
	HighRes     string
	MediumRes   string
	DefaultRes  string
	StandardRes string
}

type Snippet struct {
	ChannelId            string
	ChannelTitle         string
	Description          string
	LiveBroadcastContent string
	PublishedAt          string
	Thumbnails           Thumbnails
	Title                string
}

type Channel struct {
	ChannelId string
	Snippet   Snippet
}

type Video struct {
	VideoId string
	Snippet Snippet
}

type Playlist struct {
	PlaylistId string
	Snippet    Snippet
}

type Items struct {
	Channels  []Channel
	Playlists []Playlist
	Videos    []Video
}

type PageInfo struct {
	ResultsPerPage int64
	TotalResults   int64
}

type SearchResults struct {
	Items         Items
	NextPageToken string
	PrevPageToken string
	PageInfo      PageInfo
}
