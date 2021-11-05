package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the REST API service",
	Long:  ``,
	Run:   func(cmd *cobra.Command, args []string) { serverRun(cmd) },
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.PersistentFlags().StringP("bind", "b", viper.GetString("BIND_ADDR"), "Host or addres to bind.")
	serverCmd.PersistentFlags().IntP("port", "p", viper.GetInt("PORT"), "Listen port to use.")
	serverCmd.PersistentFlags().StringP("url", "u", viper.GetString("URL"), "Onyxia server URL")
}

func serverRun(cmd *cobra.Command) error {
	// TODO: create the HTTP REST input adaptor
	// TODO: create the output adaptor
	// TODO: create the core service

	return nil
}
