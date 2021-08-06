package cmd

import (
	"github.com/b-nova-openhub/sopagex/pkg/rest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	port   string
	solr   string
	source string

	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Serve solr-page-exposer service",
		Long:  ``,
		Run:   serve,
	}
)

func init() {
	serveCmd.PersistentFlags().StringVarP(&port, "port", "p", "8080", "This is the port that the REST endpoint is being served on. Default port if none is being specified is 8080.")
	serveCmd.PersistentFlags().StringVarP(&solr, "solr", "t", "", "This is the fully qualified URL to the endpoint of your solr instance.")
	serveCmd.PersistentFlags().StringVarP(&source, "source", "s", "", "This is a list of fully qualified URLs to your source endpoints which expose the static pages. If you got more than one entry, please separate them with a comma.")
	viper.BindPFlag("port", serveCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("solr", serveCmd.PersistentFlags().Lookup("solr"))
	viper.BindPFlag("source", serveCmd.PersistentFlags().Lookup("source"))
}

func serve(ccmd *cobra.Command, args []string) {
	rest.HandleRequests()
}
