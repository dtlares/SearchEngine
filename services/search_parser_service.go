package services

import (
	"SearchEngine/domains"
	"SearchEngine/providers"
	"sync"
)

// Search searches a query in different search engines
func Search(searchQuery string, engines []string) domains.SearchResponse {

	searchEngines := []providers.SearchProvider{}

	for _, engine := range engines {
		searchEngines = append(searchEngines, providers.CreateSearchProvider(engine))
	}

	var wg sync.WaitGroup
	searchChannel := make(chan domains.SearchResponse, len(searchEngines))

	for _, searchEngine := range searchEngines {
		wg.Add(1)
		go func() {
			searchChannel <- searchEngine.Search(searchQuery)
			wg.Done()
		}()
	}

	wg.Wait()
	close(searchChannel)

	// Merge Search results
	results := <-searchChannel
	for {
		chResult, ok := <-searchChannel
		if !ok {
			break // The channel does not have more items
		}

		for k, v := range chResult {
			results[k] = v
		}
	}

	return results
}
