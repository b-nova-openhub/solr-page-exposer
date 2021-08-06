package cmd

import (
	"github.com/b-nova-openhub/sopagex/pkg/rest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	port    string
	filter  string
	siteMap string

	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Serve solr-page-exposer service",
		Long:  ``,
		Run:   serve,
	}
)

func init() {
	serveCmd.PersistentFlags().StringVarP(&port, "port", "p", "8080", "This is the port that the REST endpoint is being served on. Default port if none is being specified is 8080.")
	viper.BindPFlag("port", serveCmd.PersistentFlags().Lookup("port"))
}

func serve(ccmd *cobra.Command, args []string) {
	rest.HandleRequests()
}
