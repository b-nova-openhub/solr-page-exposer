package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:           "sopagex",
		Short:         "sopagex – a service that sends pages to be indexed to a solr instance",
		Long:          ``,
		Version:       "0.0.0",
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	serveCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Path to custom config file. Default is '$HOME/.sopagex/config.yaml'.")
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(serveCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("/etc/sopagex/")
		viper.AddConfigPath("$HOME/.sopagex/")
		viper.AddConfigPath("../config/")
		viper.AddConfigPath(".")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Config file used for fisicus: ", viper.ConfigFileUsed())
	}
}
