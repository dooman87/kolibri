package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var (
	//If true then will log all responses using log.info logger.
	//False by default.
	LogResponse bool
)

type RequestEnhancer func(req *http.Request)

type Http struct {
	//You can set request enhancer and provide
	//additional headers/cookies for your requests
	Enhancer RequestEnhancer
}

//Sends Http GET request to url. Unmarshalling response to the target object.
//Returns error and response code.
//Error returned if non 200 code was returned or unmarshal error occurred.
func (h *Http) GetJson(urlStr string, target interface{}) (error, int) {
	if target == nil {
		return errors.New("Target interface must be specified."), -1
	}

	req := &http.Request{
		Method: "GET",
	}

	resp, err := h.executeRequest(urlStr, req)
	if err != nil {
		statusCode := -1
		if resp != nil {
			statusCode = resp.StatusCode
		}
		return err, statusCode
	}

	defer resp.Body.Close()
	jsonData, err := ioutil.ReadAll(resp.Body)

	logResponse(urlStr, resp, jsonData)

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
func (h *Http) PostJson(urlStr string, body interface{}, target interface{}) (error, int) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err, -1
	}

	req := &http.Request{
		Method: "POST",
		Header: map[string][]string{
			"Content-Type": {"application/json"},
		},
		Body: ioutil.NopCloser(bytes.NewReader(bodyBytes)),
		ContentLength: int64(len(bodyBytes)),
	}

	resp, err := h.executeRequest(urlStr, req)
	if err != nil {
		return err, -1
	}

	defer resp.Body.Close()
	jsonData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, -1
	}

	logResponse(urlStr, resp, jsonData)

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

func (h *Http) executeRequest(urlStr string, req *http.Request) (*http.Response, error) {
	if len(urlStr) == 0 {
		return nil, errors.New("Url is required")
	}

	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, fmt.Errorf("Error parsing url: %v", err)
	}
	req.URL = u
	if h.Enhancer != nil {
		h.Enhancer(req)
	}

	return http.DefaultClient.Do(req)
}

func logResponse(url string, resp *http.Response, body []byte) {
	if LogResponse {
		log.Printf("Request: [%s]\n Code: [%d]\n Response: [%s]\n", url, resp.StatusCode, string(body))
	}
}
