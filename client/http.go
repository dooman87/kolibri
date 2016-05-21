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
//Returns error if non 200 code was returned or unmarshal error occurred.
func GetJson(url string, target interface{}) error {
	if target == nil || len(url) == 0 {
		return errors.New("Both args, url and target, are required")
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	jsonData, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if LogResponse {
		log.Printf("Request: [%s]\n Code: [%d]\n Response: [%s]\n", url, resp.StatusCode, string(jsonData))
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("Expected code %d but got %d,\n response: [%s]",
			http.StatusOK, resp.StatusCode, string(jsonData)))
	}

	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, target)

	if err != nil {
		return err
	}
	return nil

}
