package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/mgrzybek/gonyxia-api/internal/core"
	"github.com/mgrzybek/gonyxia-api/internal/inputs"

	log "github.com/sirupsen/logrus"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the REST API service",
	Long:  ``,
	Run:   func(cmd *cobra.Command, args []string) { serverRun(cmd) },
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.PersistentFlags().StringP("bind-addr", "b", viper.GetString("BIND_ADDR"), "Host or addres to bind.")
	serverCmd.PersistentFlags().StringP("publish-url", "u", viper.GetString("PUBLISH_URL"), "Onyxia server URL")
	serverCmd.PersistentFlags().StringP("regions-json", "r", "", "Regions JSON file")
	serverCmd.PersistentFlags().StringP("catalogs-json", "c", "", "Catalog JSON file")
}

func serverRun(cmd *cobra.Command) error {
	/*
	 * Ops options: binding
	 */
	bindAddr, _ := cmd.Flags().GetString("bind-addr")
	publishAddr, _ := cmd.Flags().GetString("publish-addr")
	verifyHTTPServerOptions(bindAddr, publishAddr)

	/*
	 * Ops options: logs
	 */
	logLevel, _ := cmd.Flags().GetString("log-level")
	logFormat, _ := cmd.Flags().GetString("log-format")
	configureLogger(logLevel, logFormat)

	/*
	 * Business options: regions
	 */
	regionsJSON, _ := cmd.Flags().GetString("regions-json")
	regions := loadRegionsFile(regionsJSON)

	/*
	 * Business options: catalogs
	 */
	catalogsJSON, _ := cmd.Flags().GetString("catalogs-json")
	catalogs := loadCatalogsFile(catalogsJSON)

	// TODO: Create the Orchestrator adapter
	// TODO: Get k8s credentials
	//orchestrator := backoffice.Kubernetes{}

	// Create the core service
	engine, err := core.NewEngine(
		regions,
		catalogs,
	)
	if err != nil {
		log.Panic(err)
	}

	// Create the HTTP REST input adaptor
	service := inputs.NewRestService(
		bindAddr,
		publishAddr,
		&engine,
	)
	service.Run()

	return nil
}
