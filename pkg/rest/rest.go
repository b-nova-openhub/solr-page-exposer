package rest

import (
	"github.com/b-nova-openhub/sopagex/pkg/solr"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/expose", expose).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+viper.GetString("port"), router))
}

func expose(w http.ResponseWriter, r *http.Request) {
	//solr.GetAll()

	myList := map[string]string{
		"keywords":    "testKeyword",
		"author":      "testAuthor",
		"description": "testDesc",
		"title":       "testTitle",
		"url":         "testUrl1",
		"article":     "testUrl",
	}

	solr.AddValuesToSolr(myList)
}
