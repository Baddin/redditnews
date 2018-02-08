package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Item struct {
	Author string `json:"author"`
	Score  int    `json:"score"`
	URL    string `json:"url"`
	Title  string `json:"title"`
}

type response struct {
	Data1 struct {
		Children []struct {
			Data2 Item `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

func main() {
	resp, err := http.Get("https://www.reddit.com/r/golang.json")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status)
	}

	r := new(response)
	err = json.NewDecoder(resp.Body).Decode(r)

	for _, child := range r.Data1.Children {
		fmt.Println(child.Data2)
		fmt.Println(child.Data2.Author)
		fmt.Println(child.Data2.Score)
		fmt.Println(child.Data2.URL)
		fmt.Println(child.Data2.Title)
	}

}
