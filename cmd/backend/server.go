package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Jira-Analyzer/backend-services/internal/app"
	"github.com/Jira-Analyzer/backend-services/internal/config"
	"github.com/Jira-Analyzer/backend-services/internal/logger"
	log "github.com/sirupsen/logrus"
)

const (
	logFile    string = "logs.log"
	errLogFile string = "err_logs.log"
	configFile string = "configs/config.yaml"
)

func main() {
	logs, err := os.Create(logFile)
	if err != nil {
		log.Fatal(fmt.Errorf("Unable to create %s file. Error: %w", logFile, err))
	}
	defer logs.Close()

	errlogs, err := os.Create(errLogFile)
	if err != nil {
		log.Fatal(fmt.Errorf("Unable to create %s file. Error: %w", logFile, err))
	}
	defer errlogs.Close()

	logger.SetupLogrus(logs, errlogs)

	conf, err := config.ReadConfigFromYAML(configFile)
	if err != nil {
		log.Fatal(fmt.Errorf("Read of config from %s failed: %w", configFile, err))
	}
	err = config.ValidateConfig(conf)
	if err != nil {
		log.Fatal(fmt.Errorf("%s parsing failed: %w", configFile, err))
	}

	log.Info("Starting...")
	app := app.NewApp(conf)
	app.Start()

	interupt := make(chan os.Signal, 1)
	signal.Notify(interupt, os.Interrupt, syscall.SIGTERM)

	select {
	case serr := <-app.GetServerNotify():
		log.Error(fmt.Errorf("Server closes with error: %w", serr))
	case signl := <-interupt:
		log.Info("Cought signal while App running: " + signl.String())
	}

	log.Info("Shutting down...")
	app.Stop()
}
