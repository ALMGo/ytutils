package ytutils

import (
	"fmt"
	"google.golang.org/api/youtube/v3"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func GetSearchSuggestions(q string) []string {
	u, _ := url.Parse(searchSuggestionsUrl)
	values, _ := url.ParseQuery(u.RawQuery)
	values.Set("q", q)

	u.RawQuery = values.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	s := strings.Split(string(body), "[")
	var suggestions []string
	for i, s := range s {
		s2 := strings.Split(s, "\"")
		if i == 1 || len(s2) < 2 || s2[1] == "k" {
			continue
		}
		suggestions = append(suggestions, s2[1])
	}
	return suggestions
}

func ProcessSearchResults(search *youtube.SearchListResponse) SearchResults {
	var videos []Video
	var playlists []Playlist
	var channels []Channel

	for _, item := range search.Items {
		switch item.Id.Kind {
		case "youtube#video":
			videos = append(videos, processVideoItem(item))
		case "youtube#channel":
			channels = append(channels, processChannelItem(item))
		case "youtube#playlist":
			playlists = append(playlists, processPlaylistItem(item))
		}
	}
	return SearchResults{
		Items: Items{
			Videos: videos,
			Channels: channels,
			Playlists: playlists,
		},
		NextPageToken: search.NextPageToken,
		PrevPageToken: search.PrevPageToken,
		PageInfo: PageInfo{
			ResultsPerPage: search.PageInfo.ResultsPerPage,
			TotalResults: search.PageInfo.TotalResults,
		},
	}
}

func Search(service *youtube.Service, q string, options *SearchOptions) (*youtube.SearchListResponse, error) {
	finalOptions := validateSearchOptions(options)

	call := service.Search.List([]string{"id","snippet"}).
		Q(q).
		MaxResults(finalOptions.MaxResults).
		Order(finalOptions.Order).
		PageToken(finalOptions.PageToken).
		Type(finalOptions.SearchType...)
	return call.Do()
}
