package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	//If true then will log all responses using log.info logger.
	//False by default.
	LogResponse bool
)

//Sends Http GET request to url. Unmarshalling response to the target object.
//Returns error and response code.
// Error returned if non 200 code was returned or unmarshal error occurred.
func GetJson(url string, target interface{}) (error, int) {
	if target == nil || len(url) == 0 {
		return errors.New("Both args, url and target, are required"), -1
	}

	resp, err := http.Get(url)
	if err != nil {
		return err, resp.StatusCode
	}

	defer resp.Body.Close()
	jsonData, err := ioutil.ReadAll(resp.Body)

	if LogResponse {
		log.Printf("Request: [%s]\n Code: [%d]\n Response: [%s]\n", url, resp.StatusCode, string(jsonData))
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("Expected code %d but got %d,\n response: [%s]",
			http.StatusOK, resp.StatusCode, string(jsonData))), resp.StatusCode
	}

	if err != nil {
		return err, resp.StatusCode
	}

	err = json.Unmarshal(jsonData, target)

	if err != nil {
		return err, resp.StatusCode
	}
	return nil, resp.StatusCode
}
