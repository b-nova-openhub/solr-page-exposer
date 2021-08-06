package solr

import (
	"github.com/vanng822/go-solr/solr"
)

func GetAllSolrDocs(solrUrl, solrCore string) []solr.Document {
	si, _ := solr.NewSolrInterface(solrUrl, solrCore)
	query := solr.NewQuery()
	query.Q("*:*")
	s := si.Search(query)
	r, _ := s.Result(nil)
	return r.Results.Docs
}

func AddValuesToSolr(solrUrl string, blogContent map[string]string) {
	d := solr.Document{}

	for key, value := range blogContent {
		d.Set(key, value)
	}
	executeAddingToSolr(solrUrl, d)
}

func executeAddingToSolr(solrUrl string, solarDocument solr.Document) {
	si, _ := solr.NewSolrInterface(solrUrl, "blogCore")

	documentSlice := make([]solr.Document, 0, 1)
	documentSlice = append(documentSlice, solarDocument)

	si.Add(documentSlice, 11, nil)
	si.Commit()
}
