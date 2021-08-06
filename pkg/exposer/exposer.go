package exposer

import (
	"fmt"
	"github.com/b-nova-openhub/sopagex/pkg/resolver"
	"github.com/b-nova-openhub/sopagex/pkg/solr"
)

type Exposing struct {
	Success bool     `json:"success"`
	Errors  []string `json:"errors"`
}

var Exposed *Exposing

func SyncWithSolr(solrUrl string, sources []string) *Exposing {
	pages := resolver.GetAllStatisPages(sources)
	fmt.Println(pages)

	solrDocs := solr.GetAllSolrDocs(solrUrl, "blogCore")
	fmt.Println(solrDocs)

	myList := map[string]string{
		"keywords":    "testKeyword",
		"author":      "testAuthor",
		"description": "testDesc",
		"title":       "testTitle",
		"url":         "testUrl1",
		"article":     "testUrl",
	}

	solr.AddValuesToSolr(solrUrl, myList)

	Exposed = new(Exposing)
	Exposed.Success = true
	Exposed.Errors = make([]string, 0)
	return Exposed
}
