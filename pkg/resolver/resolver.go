package resolver

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type StaticPage struct {
	Title        string `json:"title"`
	Permalink    string `json:"permalink"`
	Author       string `json:"author"`
	Tags         string `json:"tags"`
	Categories   string `json:"categories"`
	PublishDate  string `json:"publishDate"`
	Description  string `json:"description"`
	ShowComments string `json:"showComments"`
	IsPublished  string `json:"isPublished"`
	Body         string `json:"body"`
}

func GetAllStatisPages(urls []string) []StaticPage {
	allStaticPages := make([]StaticPage, 0, 0)
	for _, url := range urls {
		pages := getResolvedPages(url)
		allStaticPages = append(allStaticPages, pages...)
	}
	return allStaticPages
}

func getResolvedPages(url string) []StaticPage {
	get, getErr := http.Get(url)
	if getErr != nil {
		log.Fatal(getErr)
	}
	body, readErr := ioutil.ReadAll(get.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	var staticPages []StaticPage
	jsonErr := json.Unmarshal(body, &staticPages)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return staticPages
}
