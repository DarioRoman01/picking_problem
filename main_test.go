package main

import "testing"

func TestFirstScenario(t *testing.T) {
	// const (
	// 	url       = "//www.archdaily.com/345048/nirma-vidyavihar-apurva-amin-architects"
	// 	token     = "345048"
	// 	thumbnail = "//images.adsttc.com/media/images/5143/8a8e/b3fc/4baa/2c00/000e/small_jpg/NVV1_AAA_09.jpg?1363380874"
	// 	title     = "Nirma Vidyavihar / Apurva Amin Architects"
	// )
	content, err := ParseFile("./tests/test1.json")
	switch {
	case err != nil:
		t.Error("There was an error while parsing the file")

	case len(content.Logs) != 1:
		t.Errorf("the length of the Logs array should be 1 got %d", len(content.Logs))

	case len(content.Myad) != 0:
		t.Errorf("the length of the content array should be 0 got %d", len(content.Content))

	case len(content.Content) != 0:
		t.Errorf("the length of the content array should be 0 got %d", len(content.Content))
	}
}
