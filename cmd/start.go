package cmd

import (
	"github.com/spf13/cobra"
	"proxyserver/config"
	"proxyserver/server"
)

var configPath string

func init() {
	versionCmd.Flags().StringVar(&configPath, "config", "", "config file (JSON)")
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "start",
	Short: "Where it all begins",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		config.Parse(&configPath)
		srv := server.Create()
		err := srv.ListenAndServe()
		if err != nil {
			panic(err)
		}
	},
}
