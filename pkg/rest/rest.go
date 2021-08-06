package rest

import (
	"encoding/json"
	"fmt"
	"github.com/b-nova-openhub/sopagex/pkg/exposer"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strings"
)

func HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/expose", getExpose).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+viper.GetString("port"), router))
}

func getExpose(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Expose Request: %+v\n", r)
	sources := strings.Split(viper.GetString("sources"), ",")
	exposed := exposer.SyncWithSolr(viper.GetString("solr"), sources)
	fmt.Printf("Forward Response: %+v\n", exposed)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exposed)
}
