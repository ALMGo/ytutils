package ytutils

import "os"

var developerKey = os.Getenv("GOOGLE_API_KEY")

const searchSuggestionsUrl = `https://clients1.google.com/complete/search?client=youtube&gs_ri=youtube&ds=yt`

var validOrders = map[string]bool{
	"date": true,
	"rating": true,
	"relevance": true,
	"title": true,
	"videoCount": true,
	"viewCount": true,
}

var validTypes = map[string]bool{
	"video": true,
	"playlist": true,
	"channel": true,
}

var defaultSearchOptions = SearchOptions{
	PageToken:  "",
	Order:      "relevance",
	SearchType: []string{"video"},
	MaxResults: 20,
}
