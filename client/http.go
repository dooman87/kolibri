package client

import (
	"bytes"
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
//Error returned if non 200 code was returned or unmarshal error occurred.
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

	logResponse(url, resp, jsonData)

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

//Sends Http POST request to url. Marshalling body to JSON and expecting JSON
//in response. If target is nil then will omit response.
//Returns error if non 20x code was received. Second return parameter is status code.
func PostJson(url string, body interface{}, target interface{}) (error, int) {
	if len(url) == 0 {
		return errors.New("Url is required"), -1
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err, -1
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(bodyBytes))
	if err != nil {
		return err, -1
	}

	defer resp.Body.Close()
	jsonData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, -1
	}

	logResponse(url, resp, jsonData)

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return errors.New(fmt.Sprintf("Expected code %s but got %d,\n response: [%s]",
			"20x", resp.StatusCode, string(jsonData))), resp.StatusCode
	}

	if target != nil {
		err = json.Unmarshal(jsonData, target)
		if err != nil {
			return err, resp.StatusCode
		}
	}

	return nil, resp.StatusCode
}

func logResponse(url string, resp *http.Response, body []byte) {
	if LogResponse {
		log.Printf("Request: [%s]\n Code: [%d]\n Response: [%s]\n", url, resp.StatusCode, string(body))
	}
}
