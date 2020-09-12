package main

import (
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/mvanyushkin/otus-social-network/internal/config"
	"github.com/mvanyushkin/otus-social-network/internal/domain"
	"github.com/mvanyushkin/otus-social-network/internal/http"
	"github.com/mvanyushkin/otus-social-network/logger"
	log "github.com/sirupsen/logrus"
)

func main() {
	configFilePath := flag.String("config", "", "settings file")
	flag.Parse()
	if configFilePath == nil {
		defaultConfigFileName := "settings.json"
		configFilePath = &defaultConfigFileName
	}

	cfg, err := config.GetConfig(configFilePath)
	if err != nil {
		log.Fatalf("The config file is broken: %v", err.Error())
	}

	logger.SetupLogger(cfg.LogFile, cfg.LogLevel)
	log.Info("application started.")

	db, err := sqlx.Open("mysql", cfg.ConnectionString)
	if err != nil {
		log.Fatalf("unable to open db connection %v", err.Error())
	}

	profileService := domain.NewProfileService(db)
	searcher := domain.NewProfileSearcher(db)
	accountService := domain.NewAccountService(db)
	err = http.Serve(cfg.HttpListen, profileService, searcher, accountService)
	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Info("application is shutdown")
}
