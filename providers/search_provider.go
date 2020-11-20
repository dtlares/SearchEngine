package providers

import (
	"SearchEngine/domains"
)

// SearchProvider a search Provider Interface
type SearchProvider interface {
	Search(searchString string) domains.SearchResponse
}

// CreateSearchProvider creates a Search Provider
func CreateSearchProvider(selectedProvider string) SearchProvider {
	switch selectedProvider {
	case "google":
		return GoogleProvider{}
	case "bing":
		return BingProvider{}
	}
	return GoogleProvider{}
}
