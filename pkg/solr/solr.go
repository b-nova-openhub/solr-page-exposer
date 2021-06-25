package solr

import (
	"fmt"
	"github.com/vanng822/go-solr/solr"
)

func GetAll()  {
	si, _ := solr.NewSolrInterface("http://localhost:8983/solr", "blogCore")
	query := solr.NewQuery()
	query.Q("*:*")
	s := si.Search(query)
	r, _ := s.Result(nil)
	fmt.Println(r.Results.Docs)

}

func AddValuesToSolr(blogContent map[string]string ){
	d := solr.Document{}

	for key, value := range blogContent {
		d.Set(key, value)
	}
	executeAddingToSolr(d)
}

func executeAddingToSolr(solarDocument solr.Document){
	si, _ := solr.NewSolrInterface("http://localhost:8983/solr", "blogCore")

	documentSlice := make([]solr.Document, 0, 1)
	documentSlice  = append(documentSlice, solarDocument)

	si.Add(documentSlice,11,nil)
	si.Commit()
}





