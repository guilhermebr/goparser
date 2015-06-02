package parser

import (
	"net/http"
	"testing"
)

var testJSON = []byte(`[{
	"name": "James",
	"email": "james@bond.com",
	"gender":  "male",
	"age":     "50",
    "missions": 200,
    "failures": 0
}]`)

const json_url = "http://api.reddit.com/r/golang/hot"

func TestParseJSON(t *testing.T) {
	data := ParseJSON(testJSON)

	if data[0].Name != "James" {
		t.Fatalf("expected 'James', got %v", data[0].Name)
	}

	if data[0].Email != "james@bond.com" {
		t.Fatalf("expected 'james@bond.com', got %v", data[0].Email)
	}

	if data[0].Gender != "male" {
		t.Fatalf("expected 'male', got %v", data[0].Gender)
	}

	if data[0].Age != "50" {
		t.Fatalf("expected 50, got %#v", data[0].Age)
	}

}

func TestGetContentJSON(t *testing.T) {
	resp, _ := http.Get(json_url)

	data := getContentType(resp.Header.Get("Content-Type"))

	if data != "json" {
		t.Fatalf("expected 'json', got %v", data)
	}
}
