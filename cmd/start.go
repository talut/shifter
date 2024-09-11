package cmd

import (
	"github.com/spf13/cobra"
	"proxyserver/config"
	"proxyserver/server"
)

var configPath string

func init() {
	startCMD.Flags().StringVar(&configPath, "config", "", "config file (JSON)")
	rootCmd.AddCommand(startCMD)
}

var startCMD = &cobra.Command{
	Use:   "start",
	Short: "Where it all begins",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		if configPath == "" {
			configPath = "config.json"
		}
		config.Parse(&configPath)
		srv := server.Create()
		err := srv.ListenAndServe()
		if err != nil {
			panic(err)
		}
	},
}
