package cmd

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/mgrzybek/gonyxia-api/internal/backoffice"
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

func verifyHTTPServerOptions(bindAddr, publishAddr string) {
	log.Debug("bindAddr: " + bindAddr)
	if len(bindAddr) == 0 {
		log.Fatal("The given bindAddr is invalid")
	}
}

func configureLogger(level, format string) {
	switch strings.ToLower(level) {
	case "trace":
		log.SetLevel(log.TraceLevel)
		log.SetReportCaller(true)
	case "debug":
		log.SetLevel(log.DebugLevel)
		log.SetReportCaller(true)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	default:
		log.Panic("The given log level is not recognised.")
	}

	if strings.ToLower(format) == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	}
}

func loadCatalogsFile(f string) []core.Catalog {
	data, err := ioutil.ReadFile(f)
	log.Trace("Catalog JSON data: ", string(data))
	if err != nil {
		log.Fatal("cannot load catalogs configuration: ", err)
	}

	var catalogs []core.Catalog
	err = json.Unmarshal(data, &catalogs)
	if err != nil {
		log.Fatal("cannot load catalogs configuration: ", err)
	}

	log.Trace("Catalogs from configuration: ", catalogs)
	return catalogs
}

func loadRegionsFile(f string) []core.Region {
	data, err := ioutil.ReadFile(f)
	log.Trace("Region JSON data: ", string(data))
	if err != nil {
		log.Fatal("cannot load regions configuration: ", err)
	}

	var regions []core.Region
	err = json.Unmarshal(data, &regions)
	if err != nil {
		log.Fatal("cannot load regions configuration: ", err)
	}

	log.Trace("Regions from configuration: ", regions)

	for i := range regions {
		regions[i].Services.Driver, err = backoffice.NewKubernetes(
			regions[i].Services.Server.ConfigFile,
		)

		if err != nil {
			log.Panic("Cannot create Kubernetes connection: ", err)
		}
	}

	return regions
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
