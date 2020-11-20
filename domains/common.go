package domains

// PageResult is a search result
type PageResult struct {
	Title string
	URL   string
}

// SearchResponse represents the result of a search
type SearchResponse map[string]PageResult
