package alert_manager

import (
	"context"
	"flag"
	"github.com/golang/glog"
	"github.com/mayuresh82/alert_manager/api"
	ah "github.com/mayuresh82/alert_manager/handler"
	"github.com/mayuresh82/alert_manager/internal/models"
	"github.com/mayuresh82/alert_manager/internal/stats"
	"github.com/mayuresh82/alert_manager/plugins"
	"os"
	"os/signal"
	"syscall"
)

// global flags
var (
	alertConfig = flag.String("alert-config", "", "full path to alert defintion file")
)

func Run(config *Config) {
	if config.Agent.TeamName == "" {
		config.Agent.TeamName = "default"
	}
	models.TeamName = config.Agent.TeamName
	db := models.NewDB(config.Db.Addr, config.Db.Username, config.Db.Password, config.Db.DbName, config.Db.Timeout)
	defer db.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// start the config loader
	reloadConfig := make(chan struct{})
	ah.Config = ah.NewConfigHandler(*alertConfig)
	go func() {
		for {
			select {
			case <-reloadConfig:
				ah.Config.LoadConfig()
			case <-ctx.Done():
				return
			}
		}
	}()

	// start the handler
	handler := ah.NewHandler(db)
	go handler.Start(ctx)

	//Initialize all the plugins
	// Listener, transforms
	plugins.Init(ctx, db)

	// start the API server
	glog.Infof("Starting API server on %s", config.Agent.ApiAddr)
	server := api.NewServer(config.Agent.ApiAddr, handler)
	go server.Start(ctx)

	// start the reporting agent
	glog.Infof("Will send stats to %s", config.Reporter.Url)
	go stats.StartExport(ctx, config.Agent.StatsExportInterval)
	go config.Reporter.Start(ctx)

	// wait for sig
	signalChan := make(chan os.Signal, 1)
	shutdown := make(chan struct{})
	signal.Notify(signalChan, os.Interrupt, syscall.SIGHUP, syscall.SIGTERM)
	go func() {
		for {
			sig := <-signalChan
			if sig == os.Interrupt || sig == syscall.SIGTERM {
				glog.Infof("Alert Manager shutting down")
				shutdown <- struct{}{}
				return
			}
			if sig == syscall.SIGHUP {
				glog.Infof("Reloading alert config")
				reloadConfig <- struct{}{}
				// TODO restart the processors ?
			}
		}
	}()
	<-shutdown
}
