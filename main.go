package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

// Read the json file in the given path
func ReadFile(path string) []byte {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("It was not posible to open the File: %s\n", err.Error())
	}

	return content
}

// decode the content of the json file in the recommendations struct
func parseJson(content []byte) *Recommendations {
	var payload Recommendations
	err := json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatalf("Unable to marshal the json file %s", err.Error())
	}

	return &payload
}

func main() {
	content := ReadFile("./tests/test1.json")
	recommendations := parseJson(content)
	if recommendations.Logs != nil {
		fmt.Println(recommendations.Logs)
	}

	if len(recommendations.Content) != 0 {
		fmt.Println("Fallo")
	}

	if len(recommendations.Myad) != 0 {
		fmt.Println("Fallo")
	}
}
