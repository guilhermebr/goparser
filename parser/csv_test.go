package parser

import (
	"net/http"
	"strings"
	"testing"
)

const csv_url = "https://extranet.who.int/tme/generateCSV.asp?ds=dictionary"

var testCSV = "name,email,age,country\nJames,james@bond.com,50,UK\nChuck,Norris,60,USA"

func TestParseCSV(t *testing.T) {
	var reader = strings.NewReader(testCSV)

	data, _ := ParseCSV(reader)

	if data[0].Name != "James" {
		t.Fatalf("expected 'James', got %v", data[0].Name)
	}

	if data[0].Email != "james@bond.com" {
		t.Fatalf("expected 'james@bond.com', got %v", data[0].Email)
	}

	if data[0].Age != "50" {
		t.Fatalf("expected 50, got %#v", data[0].Age)
	}
}

func TestGetContentCSV(t *testing.T) {
	resp, _ := http.Get(csv_url)

	data := getContentType(resp.Header.Get("Content-Type"))

	if data != "csv" {
		t.Fatalf("expected 'csv', got %v", data)
	}
}
