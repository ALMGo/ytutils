package ytutils

import "google.golang.org/api/youtube/v3"

func processThumbnails(thumbs *youtube.ThumbnailDetails) Thumbnails {
	return Thumbnails{
		DefaultRes:  thumbs.Default.Url,
		MediumRes:   thumbs.Medium.Url,
		MaxRes:      thumbs.Maxres.Url,
		HighRes:     thumbs.High.Url,
		StandardRes: thumbs.Standard.Url,
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
