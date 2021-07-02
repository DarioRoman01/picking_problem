package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"github.com/Haizza1/picking_problem/leveshtein"
)

// util struct to represent a tuple of article token and title
type Tuple struct {
	Token string
	Title string
}

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

// util function to push item in the slice concurrently
func push(titles []*Tuple, articles []Article, wg *sync.WaitGroup) []*Tuple {
	for _, article := range articles {
		titles = append(titles, &Tuple{article.Token, article.Title})
	}

	wg.Done()
	return titles
}

func FindRecommendations(recommendations *Recommendations) []string {
	var wg sync.WaitGroup
	var titles []*Tuple

	wg.Add(3)

	go func() {
		titles = push(titles, recommendations.Logs, &wg)
	}()

	go func() {
		titles = push(titles, recommendations.Content, &wg)
	}()

	go func() {
		titles = push(titles, recommendations.Myad, &wg)
	}()

	wg.Wait()
	return sort(titles)

}

func sort(arr []*Tuple) []string {
	var final []string
	var sorted []*Tuple
	if len(arr) <= 4 {
		for i := 0; i < len(arr); i++ {
			final = append(final, arr[i].Token)
		}

		return final
	}

	sorted = append(sorted, arr[0])
	fmt.Println(arr[0].Token)
	toCompare := arr[0]
	lastDist := 10000000
	lastIndex := 0

	for i := 1; i < len(arr); i++ {
		dist := leveshtein.Dist(arr[i].Title, toCompare.Title)
		if dist < lastDist {
			sorted = Insert(sorted, lastIndex+1, arr[i])
			lastIndex += 1
			lastDist = dist
		} else {
			sorted = append(sorted, arr[i])
		}

	}

	for i := 0; i < 4; i++ {
		final = append(final, sorted[i].Token)
	}

	return final
}

// util function to insert values in the array at given position
func Insert(a []*Tuple, index int, value *Tuple) []*Tuple {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}

	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func main() {
	recommendations, err := ParseFile("./tests/test4.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Testing seceneario 4....")
	result := FindRecommendations(recommendations)
	fmt.Println("The result is: ", result)
}
