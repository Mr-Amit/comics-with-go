package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type Result struct {
	Month     string `json:"month"`
	Num       int    `json:"num"`
	Link      string `json:"link"`
	Year      string `json:"year"`
	News      string `json:"News"`
	SafeTitle string `json:"safetitle"`
	Trascript string `json:"trascript"`
	Alt       string `json:"alt"`
	Img       string `json:"img"`
	Title     string `json:"title"`
	Day       string `json:"day"`
}

const Url = "https://xkcd.com"

func fetch(n int) (*Result, error) {

	client := &http.Client{
		Timeout: 5 * time.Minute,
	}

	url := strings.Join([]string{Url, fmt.Sprintf("%d", n), "info.0.json"}, "/")

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("http request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http err: %v", err)
	}

	var data Result

	if resp.StatusCode != http.StatusOK {
		data = Result{}
	} else {
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return nil, fmt.Errorf("json err: %v", err)
		}
	}

	resp.Body.Close()

	return &data, nil
}

func main() {
	n := 200
	result, err := fetch(n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", result)
}
