package main

import (
	"CodeChallenge/DriversMetricsImporter/utils/config"
	"CodeChallenge/DriversMetricsImporter/utils/net/http"
	"CodeChallenge/DriversMetricsImporter/utils/log"
	"CodeChallenge/DriversMetricsImporter/data-providers"
	"CodeChallenge/DriversMetricsImporter/models"
)

/*
	Holds app config key-value data
*/
var appConfig map[string]interface{};

/*
	Initializer for configuration
 */
func init() {
	appConfig = config.GetAppConfiguration()
}

/*
	Imports drivers data from URL and save them on DB
*/
func importExportDrivers() error {
	var driversURL string = appConfig["driversURL"].(string);
	var drivers []models.Driver
	var err = http.ImportJSON(driversURL, &drivers, 0);
	if (err != nil) {
		return err
	}

	err = data_providers.ExportDrivers(drivers);
	if (err != nil) {
		return err
	}

	return nil
}

/*
	Imports metrics data from URL and save them on DB
*/
func importExportMetrics() error {
	var metrics []models.Metric
	var metricsURL string = appConfig["metricsURL"].(string);

	err := http.ImportJSON(metricsURL, &metrics, '\n');
	if (err != nil ) {
		return err
	}

	data_providers.ExportMetrics(metrics)

	return nil
}

/*
	Main routine - Imports and exports drivers info and dependant metrics
*/
func main() {
	log.Info(" ------ Starting DriversMetricsImporter on CodeChallange... -----")
	var err = importExportDrivers()
	if (err != nil) {
		panic(err)
	}

	err = importExportMetrics()
	if (err != nil) {
		panic(err)
	}

	log.Info(" ------ DriversMetricsImporter on CodeChallange Ended -----")
}
