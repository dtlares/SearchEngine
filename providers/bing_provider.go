package providers

import (
	"SearchEngine/domains"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// BingProvider links to Bing
type BingProvider struct {
}

// The is the struct for the data returned by Bing.
type bingAnswer struct {
	Type         string `json:"_type"`
	QueryContext struct {
		OriginalQuery string `json:"originalQuery"`
	} `json:"queryContext"`
	WebPages struct {
		WebSearchURL          string `json:"webSearchUrl"`
		TotalEstimatedMatches int    `json:"totalEstimatedMatches"`
		Value                 []struct {
			ID               string    `json:"id"`
			Name             string    `json:"name"`
			URL              string    `json:"url"`
			IsFamilyFriendly bool      `json:"isFamilyFriendly"`
			DisplayURL       string    `json:"displayUrl"`
			Snippet          string    `json:"snippet"`
			DateLastCrawled  time.Time `json:"dateLastCrawled"`
			SearchTags       []struct {
				Name    string `json:"name"`
				Content string `json:"content"`
			} `json:"searchTags,omitempty"`
			About []struct {
				Name string `json:"name"`
			} `json:"about,omitempty"`
		} `json:"value"`
	} `json:"webPages"`
	RelatedSearches struct {
		ID    string `json:"id"`
		Value []struct {
			Text         string `json:"text"`
			DisplayText  string `json:"displayText"`
			WebSearchURL string `json:"webSearchUrl"`
		} `json:"value"`
	} `json:"relatedSearches"`
	RankingResponse struct {
		Mainline struct {
			Items []struct {
				AnswerType  string `json:"answerType"`
				ResultIndex int    `json:"resultIndex"`
				Value       struct {
					ID string `json:"id"`
				} `json:"value"`
			} `json:"items"`
		} `json:"mainline"`
		Sidebar struct {
			Items []struct {
				AnswerType string `json:"answerType"`
				Value      struct {
					ID string `json:"id"`
				} `json:"value"`
			} `json:"items"`
		} `json:"sidebar"`
	} `json:"rankingResponse"`
}

// Search searches in Bing
func (p BingProvider) Search(searchString string) domains.SearchResponse {

	const endpoint = "https://api.cognitive.microsoft.com/bing/v7.0/search"
	token := "" // os.Getenv("SECRET_BING_KEY")

	// Declare a new GET request.
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		panic(err)
	}

	// Add the payload to the request.
	param := req.URL.Query()
	param.Add("q", searchString)
	req.URL.RawQuery = param.Encode()

	// Insert the request header.
	req.Header.Add("Ocp-Apim-Subscription-Key", token)

	// Instantiate a client.
	client := new(http.Client)

	// Send the request to Bing.
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// Close the connection.
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Create a new answer.
	ans := new(bingAnswer)
	err = json.Unmarshal(body, &ans)
	if err != nil {
		fmt.Println(err)
	}

	results := make(domains.SearchResponse)

	for _, result := range ans.WebPages.Value {
		results[result.Name] = domains.PageResult{Title: result.Name, URL: result.URL}
	}
	return results
}
