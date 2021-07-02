package main

import (
	"encoding/json"
	"io/ioutil"
)

// Represents the article object
type Article struct {
	Url       string `json:"url"`       // represents the url of the article
	Token     string `json:"token"`     // represents the unique token of the article
	Thumbnail string `json:"thumbnail"` // represents the image thumbnail
	Title     string `json:"title"`     // represents the Article title
}

// Read the json file in the given path and decode
// the content of the json file in the recommendations struct
func ParseFile(path string) (map[string][]Article, error) {
	// we read the file with read file function
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var payload map[string][]Article
	json.Unmarshal(content, &payload)
	return payload, nil
}

// return the total length of the recommended articles
func totalLen(r map[string][]Article) int {
	return len(r["logs"]) + len(r["content"]) + len(r["myad"])
}

// Find most homogenues recommendatios based on the Road Robin aproach
func FindRecommendations(recommendations map[string][]Article) []string {
	var tokens []string
	totallen := totalLen(recommendations)
	index := 0

	// loop until we find the most homogenues list of recommendations.
	// The worst case is have recommendations in only one algorithm.
	for i := 0; i < 4; i++ {
		if len(tokens) == totallen {
			return tokens
		}

		for _, value := range recommendations {
			if len(value) == 0 || index >= len(value) {
				continue
			}

			tokens = append(tokens, value[index].Token)
			if len(tokens) >= 4 {
				return tokens
			}
		}

		index++
	}

	return tokens
}
