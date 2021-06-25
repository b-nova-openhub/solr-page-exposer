package main

import (
	"b-nova-solr-page-exposer/pkg/solr"
)

func main() {
	//solr.GetAll()

	myList := map[string]string{
		"keywords":"testKeyword",
		"author":"testAuthor",
		"description":"testDesc",
		"title":"testTitle",
		"url":"testUrl1",
		"article":"testUrl",
	}

	solr.AddValuesToSolr(myList)

}
