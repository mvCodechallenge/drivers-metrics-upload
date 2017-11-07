package models

import (
	"strconv"
	"fmt"
)

/*
	Holds Metric model
*/
type Metric struct {
	Name      string  `json:"metric_name, string, omitempty"`
	Val     string	`json:"value, string, omitempty"`
	Longitude float64	`json:"lon, number, omitempty"`
	Timestamp uint64  `json:"timestamp, number, omitempty"`
	Latitude  float64	`json:"lat, number, omitempty"`
	DrvId  string    `json:"driver_id, string, omitempty"`
	DriverId uint
	Value    float64
}

/*
	Nicely gets metric model data
*/
func (current *Metric) ToString() string {
    return fmt.Sprintf("Name: %s, Value: %f, DriverId: %d, Timestamp: %d", current.Name, current.Value, current.DriverId, current.Timestamp);
}

/*
	Adjust driver id and metric data represented as string on imported payload to numeric
*/
func (current *Metric) Adjust() error {
	drvId, err := strconv.ParseUint(current.DrvId, 10, 32);
	if (err != nil) {
		return err;
	}

	current.DriverId = uint(drvId);
	val, err := strconv.ParseFloat(current.Val, 64);
	if (err != nil) {
		return err;
	}

	current.Value = val;
	return nil;
}