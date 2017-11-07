package data_providers

import (
	"database/sql"
	_ "github.com/lib/pq"
	"drivers-metrics-upload/utils/config"
	"drivers-metrics-upload/utils/log"
	"drivers-metrics-upload/models"
	"fmt"
)

var (
	// Holds app config store
	appConfig map[string]interface{};
	// Holds DB instance
	_db *sql.DB;
)

/*
	Initialize DB instance and test it
*/
func init() {
	appConfig = config.GetAppConfiguration();
	dbConnectionString := appConfig["dbConnectionString"].(string);
	db, err := sql.Open("postgres", dbConnectionString)
	if (err != nil) {
		panic(err)
	}

	// Checks if DB is alive
	err = db.Ping()
	if (err != nil) {
		panic(err)
	}

	_db = db
}

/*
	Exports drivers to DB
*/
func ExportDrivers(drivers []models.Driver) error {
	total := len(drivers)
	if (total < 1) {
		return nil
	}

	log.Info(fmt.Sprintf("Found %d drivers to export...", total));

	var err = _db.Ping()
	if (err != nil) {
		return err
	}

	ok := 0;
	for _, currDriver := range drivers {
		result, err := _db.Query("select insertDriver($1, $2, $3)", currDriver.Id, currDriver.Name, currDriver.LicenseNumber)
		if (err != nil) {
			log.Warn(fmt.Sprintf("Couldn't add driver: <%s>\n\terror: %s", currDriver.ToString(), err.Error()))
		} else {
			result.Close();
			ok += 1;
			log.Info(fmt.Sprintf("%d/%d (%3.2f) drivers inserted...", ok, total, float32(ok)/float32(total) * 100));
		}
	}


	return nil
}

/*
	Exports metrics to DB
*/
func ExportMetrics(metrics []models.Metric) error {
	total := len(metrics)
	if (total < 1) {
		return nil
	}

	log.Info(fmt.Sprintf("Found %d metrics to export...", total));

	var err = _db.Ping()
	if (err != nil) {
		return err
	}

	ok := 0
	for _, currMetric := range metrics {
		currMetric.Adjust()
		result, err := _db.Query("select insertMetric($1, $2, $3, $4, $5, $6)", currMetric.DriverId, currMetric.Name, currMetric.Timestamp, currMetric.Longitude, currMetric.Latitude, currMetric.Value)
		if (err != nil) {
			log.Warn(fmt.Sprintf("Couldn't add metric: <%s>\n\terror: %s", currMetric.ToString(), err.Error()))
		} else {
			result.Close();
			ok += 1;
			log.Info(fmt.Sprintf("%d/%d (%3.2f) metrics inserted...", ok, total, float32(ok)/float32(total) * 100));
		}
	}

	return nil
}
