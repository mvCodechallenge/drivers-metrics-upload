package http

import (
	"encoding/json"
	"reflect"
	"net/http"
	"bufio"
	"time"
	"drivers-metrics-upload/utils/config"
	"drivers-metrics-upload/utils/log"
)

var (
	// Holds app config key-value data
	appConfig map[string]interface{};
	// Holds http client instance
	httpClient http.Client;
)

/*
	Initializes config and http client instances
*/
func init() {
	appConfig = config.GetAppConfiguration();
	requestTimeout := time.Duration(appConfig["httpRequestTimeout"].(float64));
	httpClient = http.Client{Timeout: time.Millisecond * requestTimeout}
}

/*
	Import JSON on HTTP url to a specific struct
	delimiter is needed for if the payload is a JSON data segments separated by the delimiter
*/
func ImportJSON(url string, target interface{}, delimiter byte) error {
	// Get HTTP response from URL
	response, err := httpClient.Get(url)
	if (err != nil) {
		return err
	}

	defer response.Body.Close()

	// If the payload is fully JSON convert it to the specific struct and return
	if (delimiter == 0) {
		return json.NewDecoder(response.Body).Decode(target)
	}

	/*
		If payload is a JSON data segments separated by the delimiter
		Need to assemble an array of structs type given
	*/

	reader := bufio.NewReader(response.Body)
	slice := reflect.ValueOf(target).Elem()
	typeOfSlice := slice.Type()
	slice.Set(reflect.MakeSlice(typeOfSlice, 0, 1))
	ptrToTarget := reflect.New(typeOfSlice.Elem())

	// Loop through the payload segments try to convert them to struct and add them to array
	for {
		part, err := reader.ReadBytes(delimiter);
		if (err != nil) {
			errMsg := err.Error();
			if (errMsg != "EOF") {
				log.Error(err.Error())
			}

			return nil
		}

		if (len(part) == 0) {
			return nil
		}

		if err := json.Unmarshal(part, ptrToTarget.Interface()); (err != nil) {
			return err
		}

		slice.Set(reflect.Append(slice, ptrToTarget.Elem()))
	}

	return nil
}

