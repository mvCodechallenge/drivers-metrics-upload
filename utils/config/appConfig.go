/*
	Key-Value storage handling for app config file
*/
package config

import (
	"os"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

//ToDo: add file watcher here for hot changes...

// Holds single instance of app config key-value store
var appConfig map[string] interface{};

/*
	Prints nicely the app config data on first initialize store
*/
func printConfig() {
	fmt.Println("\n--- Application configuration ---")
	for key, value := range appConfig {
		fmt.Println(fmt.Sprintf("\tkey: %s, value: %v", key, value));
	}
	fmt.Println("---------------------------------")
}

/*
	Method to retrieve and lazy initialize app config store
*/
func GetAppConfiguration() map[string] interface{} {
	if (appConfig != nil) {
		return appConfig;
	}

	dir, _ := os.Getwd();
	plan, _ := ioutil.ReadFile(dir + "/conf/config.json") // filename is the JSON file to read
	var data map[string] interface{}
	err := json.Unmarshal(plan, &data)
	if (err != nil) {
		panic(err)
	}

	appConfig = data;
	printConfig();
	return data;
}
