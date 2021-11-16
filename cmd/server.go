package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/mgrzybek/gonyxia-api/internal/core"
	"github.com/mgrzybek/gonyxia-api/internal/inputs"
	//"github.com/mgrzybek/gonyxia-api/internal/backoffice"

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

func serverRun(cmd *cobra.Command) error {
	bind_addr, _ := cmd.Flags().GetString("bind-addr")
	publish_addr, _ := cmd.Flags().GetString("publish-addr")
	log_level, _ := cmd.Flags().GetString("log-level")
	log_format, _ := cmd.Flags().GetString("log-format")

	configureLogger(log_level, log_format)
	verifyHttpServerOptions(bind_addr, publish_addr)

	// TODO: Create the Orchestrator adapter
	// TODO: Get k8s credentials
	//orchestrator := backoffice.Kubernetes{}

	// Create the core service
	engine := core.Engine{}
	//		orchestrator: orchestrator,
	//	}

	// Create the HTTP REST input adaptor
	service := inputs.RestService{
		Bind_addr:    bind_addr,
		Publish_addr: publish_addr,
		Engine:       engine,
	}
	service.Run()

	return nil
}
