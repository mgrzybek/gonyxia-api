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

func verifyHttpServerOptions(bind_addr, publish_addr string) {
	log.Debug("bind_addr: " + bind_addr)
	if len(bind_addr) == 0 {
		log.Fatal("The given bind_addr is invalid")
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
	log.Trace("Catalog JSON data: %s", string(data))
	if err != nil {
		log.Fatal("cannot load catalogs configuration: ", err)
	}

	var catalogs []core.Catalog
	err = json.Unmarshal(data, &catalogs)
	if err != nil {
		log.Fatal("cannot load catalogs configuration: ", err)
	}

	log.Trace("Catalogs from configuration: %#v", catalogs)
	return catalogs
}

func loadRegionsFile(f string) []core.Region {
	data, err := ioutil.ReadFile(f)
	log.Trace("Region JSON data: %s", string(data))
	if err != nil {
		log.Fatal("cannot load regions configuration: ", err)
	}

	var regions []core.Region
	err = json.Unmarshal(data, &regions)
	if err != nil {
		log.Fatal("cannot load regions configuration: ", err)
	}

	log.Trace("Regions from configuration: %#v", regions)

	for i, _ := range regions {
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
	bind_addr, _ := cmd.Flags().GetString("bind-addr")
	publish_addr, _ := cmd.Flags().GetString("publish-addr")
	verifyHttpServerOptions(bind_addr, publish_addr)

	/*
	 * Ops options: logs
	 */
	log_level, _ := cmd.Flags().GetString("log-level")
	log_format, _ := cmd.Flags().GetString("log-format")
	configureLogger(log_level, log_format)

	/*
	 * Business options: regions
	 */
	regions_json, _ := cmd.Flags().GetString("regions-json")
	regions := loadRegionsFile(regions_json)

	/*
	 * Business options: catalogs
	 */
	catalogs_json, _ := cmd.Flags().GetString("catalogs-json")
	catalogs := loadCatalogsFile(catalogs_json)

	// TODO: Create the Orchestrator adapter
	// TODO: Get k8s credentials
	//orchestrator := backoffice.Kubernetes{}

	// Create the core service
	engine := core.NewEngine(
		regions,
		catalogs,
	)

	// Create the HTTP REST input adaptor
	service := inputs.NewRestService(
		bind_addr,
		publish_addr,
		&engine,
	)
	service.Run()

	return nil
}
