package models

import "fmt"

/*
	Holds Driver model
*/
type Driver struct {
	Id       	    uint	`json:"id, number, omitempty"`
	Name     	    string  `json:"name, string, omitempty"`
	LicenseNumber	string	`json:"license_number, string, omitempty"`
}

/*
	Nicely gets driver model data
*/
func (current *Driver) ToString() string {
	return fmt.Sprintf("Id: %d, Name: %s, LicenseNumber: %s", current.Id, current.Name, current.LicenseNumber);
}