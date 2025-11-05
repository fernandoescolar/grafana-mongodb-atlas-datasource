package main

import (
	"os"

	"github.com/fernandoescolar/grafana-mongodb-atlas-datasource/pkg/plugin"
	"github.com/grafana/grafana-plugin-sdk-go/backend/datasource"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

func main() {
	err := datasource.Manage("fernandoescolar-grafana-mongodb-atlas-datasource", plugin.NewApp, datasource.ManageOpts{})
	if err != nil {
		log.DefaultLogger.Error(err.Error())
		os.Exit(1)
	}
}
