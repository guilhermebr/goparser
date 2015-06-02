package parser

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func getContentType(content string) string {
	fmt.Println(content)
	ct := strings.Split(content, ";")[0]

	if ct == "application/json" {
		return "json"
	}

	if ct == "text/csv" {
		return "csv"
	}

	return ""
}

func Parse(url string) error {
	// Load URL
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error:", err)
	}
	// Get Content
	content := getContentType(resp.Header.Get("Content-Type"))

	if content == "csv" {
		ParseCSV(resp.Body)
	} else if content == "json" {
		body, _ := ioutil.ReadAll(resp.Body)
		ParseJSON(body)
	} else {
		return errors.New("content-type not supported")
	}

	return nil
}
