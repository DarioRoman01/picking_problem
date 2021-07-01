package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Represents the article object
type Article struct {
	Url       string `json:"url"`       // represents the url of the article
	Token     string `json:"token"`     // represents the unique token of the article
	Thumbnail string `json:"thumbnail"` // represents the image thumbnail
	Title     string `json:"title"`     // represents the Article title
}

// Represents the reconmendations object
type Recommendations struct {
	Logs    []Article `json:"logs"`    // represents the articles catch by the log algorithm
	Content []Article `json:"content"` // represents the articles catch by the content algorithm
	Myad    []Article `json:"myad"`    // represents the articles catch by the myad algorithm
}

// Read the json file in the given path and decode
// the content of the json file in the recommendations struct
func ParseFile(path string) (*Recommendations, error) {
	// we read the file with read file function
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// we deocde the json content in
	// the recommendatios struct
	var payload Recommendations
	err = json.Unmarshal(content, &payload)
	if err != nil {
		return nil, err
	}

	// we return a pointer to the payload
	// to avoid duplicates in memory
	return &payload, nil
}

func main() {
	recommendations, err := ParseFile("./tests/test1.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(recommendations)
}
