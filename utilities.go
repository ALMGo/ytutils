package ytutils

import "google.golang.org/api/youtube/v3"

func processThumbnails(thumbs *youtube.ThumbnailDetails) Thumbnails {
	var defaultRes = ""
	if thumbs.Default != nil {
		defaultRes = thumbs.Default.Url
	}

	var mediumRes = ""
	if thumbs.Medium != nil {
		mediumRes = thumbs.Medium.Url
	}

	var maxRes = ""
	if thumbs.Maxres != nil {
		maxRes = thumbs.Maxres.Url
	}

	var highRes = ""
	if thumbs.High != nil {
		highRes = thumbs.High.Url
	}

	var standardRes = ""
	if thumbs.Standard != nil {
		standardRes = thumbs.Standard.Url
	}

	return Thumbnails{
		DefaultRes:  defaultRes,
		MediumRes:   mediumRes,
		MaxRes:      maxRes,
		HighRes:     highRes,
		StandardRes: standardRes,
	}
}

func processSnippet(snippet *youtube.SearchResultSnippet) Snippet {
	return Snippet{
		ChannelId: snippet.ChannelId,
		ChannelTitle: snippet.ChannelTitle,
		Description: snippet.Description,
		LiveBroadcastContent: snippet.LiveBroadcastContent,
		PublishedAt: snippet.PublishedAt,
		Thumbnails: processThumbnails(snippet.Thumbnails),
	}
}

func processVideoItem(item *youtube.SearchResult) Video {
	return Video{
		VideoId: item.Id.VideoId,
		Snippet: processSnippet(item.Snippet),
	}
}

func processChannelItem(item *youtube.SearchResult) Channel {
	return Channel{
		ChannelId: item.Id.ChannelId,
		Snippet: processSnippet(item.Snippet),
	}
}

func processPlaylistItem(item *youtube.SearchResult) Playlist {
	return Playlist{
		PlaylistId: item.Id.PlaylistId,
		Snippet: processSnippet(item.Snippet),
	}
}

func validateSearchOptions(options *SearchOptions) SearchOptions {
	var searchOrder = options.Order
	if !validOrders[searchOrder] {
		searchOrder = defaultSearchOptions.Order
	}

	var searchType = options.SearchType
	for _, item := range searchType{
		if !validTypes[item] {
			searchType = defaultSearchOptions.SearchType
			break
		}
	}

	var maxResults = options.MaxResults
	if maxResults <= 0 {
		maxResults = defaultSearchOptions.MaxResults
	}

	return SearchOptions{
		Order:      searchOrder,
		SearchType: searchType,
		MaxResults: maxResults,
		PageToken:  options.PageToken,
	}
}
