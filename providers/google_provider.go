package providers

import (
	"SearchEngine/domains"
	"log"
	"net/http"
	"os"

	"google.golang.org/api/customsearch/v1"
	"google.golang.org/api/googleapi/transport"
)

// GoogleProvider links to google
type GoogleProvider struct {
}

// Search returns a search in google
func (p GoogleProvider) Search(searchString string) domains.SearchResponse {

	apiKey := os.Getenv("SECRET_GOOGLE_KEY")
	cx := os.Getenv("SECRET_GOOGLE_CX")
	query := searchString

	client := &http.Client{Transport: &transport.APIKey{Key: apiKey}}

	svc, err := customsearch.New(client)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := svc.Cse.List().Cx(cx).Q(query).Do()
	if err != nil {
		log.Fatal(err)
	}

	results := make(domains.SearchResponse)
	for _, result := range resp.Items {
		results[result.Title] = domains.PageResult{Title: result.Title, URL: result.Link}
	}
	return results
}
