package picking_problem

import (
	"reflect"
	"testing"
)

func TestFirstScenario(t *testing.T) {
	content, err := parseFile("./tests/test1.json")
	switch {
	case err != nil:
		t.Error("There was an error while parsing the file")

	case len(content["logs"]) != 1:
		t.Errorf("the length of the Logs array should be 1 got %d", len(content["logs"]))

	case len(content["myad"]) != 0:
		t.Errorf("the length of the content array should be 0 got %d", len(content["myad"]))

	case len(content["content"]) != 0:
		t.Errorf("the length of the content array should be 0 got %d", len(content["content"]))
	}

	expectedTokens := []string{"345048"}
	tokens := balance(content)
	if !reflect.DeepEqual(expectedTokens, tokens) {
		t.Log("Expected tokens and actual tokens are not equal")
		t.Error("expected: \n", expectedTokens, "found: ", tokens)
	}
}

func TestSecondEscenario(t *testing.T) {
	content, err := parseFile("./tests/test2.json")
	switch {
	case err != nil:
		t.Error("There was an error while parsing the file")

	case len(content["logs"]) != 3:
		t.Errorf("the length of the Logs array should be 3 got %d", len(content["logs"]))

	case len(content["myad"]) != 3:
		t.Errorf("the length of the content array should be 3 got %d", len(content["myad"]))

	case len(content["content"]) != 4:
		t.Errorf("the length of the content array should be 4 got %d", len(content["content"]))
	}

	expectedTokens := []string{"790952", "103678", "788138", "802358"}
	tokens := balance(content)
	if !reflect.DeepEqual(expectedTokens, tokens) {
		t.Log("Expected tokens and actual tokens are not equal")
		t.Error("expected: \n", expectedTokens, "found: ", tokens)
	}
}
func TestThirdEscenario(t *testing.T) {
	content, err := parseFile("./tests/test3.json")
	switch {
	case err != nil:
		t.Error("There was an error while parsing the file")

	case len(content["logs"]) != 3:
		t.Errorf("the length of the Logs array should be 1 got %d", len(content["logs"]))

	case len(content["myad"]) != 0:
		t.Errorf("the length of the content array should be 0 got %d", len(content["myad"]))

	case len(content["content"]) != 1:
		t.Errorf("the length of the content array should be 3 got %d", len(content["content"]))
	}

	expectedTokens := []string{"103678", "790952", "802358", "788138"}
	tokens := balance(content)
	if !reflect.DeepEqual(expectedTokens, tokens) {
		t.Log("Expected tokens and actual tokens are not equal")
		t.Error("expected: \n", expectedTokens, "found: ", tokens)
	}
}

func TestFourthEscenario(t *testing.T) {
	content, err := parseFile("./tests/test4.json")
	switch {
	case err != nil:
		t.Error("There was an error while parsing the file")

	case len(content["logs"]) != 3:
		t.Errorf("the length of the Logs array should be 3 got %d", len(content["logs"]))

	case len(content["myad"]) != 0:
		t.Errorf("the length of the content array should be 0 got %d", len(content["myad"]))

	case len(content["content"]) != 4:
		t.Errorf("the length of the content array should be 4 got %d", len(content["content"]))
	}

	expectedTokens := []string{"790952", "103678", "802358", "562873"}
	tokens := balance(content)
	if !reflect.DeepEqual(expectedTokens, tokens) {
		t.Log("Expected tokens and actual tokens are not equal")
		t.Error("expected: \n", expectedTokens, "found: ", tokens)
	}
}
