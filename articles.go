package picking_problem

import (
	"encoding/json"
	"io/ioutil"
)

// Represents the article object
type article struct {
	Url       string `json:"url"`       // represents the url of the article
	Token     string `json:"token"`     // represents the unique token of the article
	Thumbnail string `json:"thumbnail"` // represents the image thumbnail
	Title     string `json:"title"`     // represents the Article title
}

// Read the json file in the given path and decode
// the content of the json file in the recommendations struct
func parseFile(path string) (map[string][]article, error) {
	// we read the file with read file function
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var payload map[string][]article
	err = json.Unmarshal(content, &payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

// return the total length of the recommended articles
func totalLen(r map[string][]article) int {
	return len(r["logs"]) + len(r["content"]) + len(r["myad"])
}

// Find most homegenous recommendatios based on the Road Robin aproach
func balance(recommendations map[string][]article) []string {
	var tokens []string
	totallen := totalLen(recommendations)

	// loop until we find the most homegenous list of recommendations.
	// The worst case is have recommendations in only one algorithm.
	for i := 0; i < 4; i++ {
		if len(tokens) == totallen {
			return tokens
		}

		for _, value := range recommendations {
			if len(value) == 0 || i >= len(value) {
				continue
			}

			tokens = append(tokens, value[i].Token)
			if len(tokens) >= 4 {
				return tokens
			}
		}

	}

	return tokens
}

// Find the most Homogenous recommendation in the json file
func FindMostHomogenousRecommendations(path string) ([]string, error) {
	recommendations, err := parseFile(path)
	if err != nil {
		return nil, err
	}

	return balance(recommendations), nil
}
