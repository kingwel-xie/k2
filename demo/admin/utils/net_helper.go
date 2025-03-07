package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/kingwel-xie/k2/core/logger"
)

var log = logger.Logger("helper")

var client = &http.Client{Timeout: 30 * time.Second}

// JsonResponse returns a json reponse from remote server
func JsonResponse(method string, url string, headers map[string]string, body []byte) ([]byte, error) {

	log.Infof("Do request: method=%v, url=%v, body=%v", method, url, string(body))

	request, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		log.Errorf("New json request err: %v", err)
		return nil, err
	}

	for k, v := range headers {
		request.Header.Add(k, v)
	}
	// the Content type is not json when globavend
	if request.Header.Get("Content-Type") == "" {
		request.Header.Add("Content-Type", "application/json")
	}
	request.Header.Add("Accept", "application/json")

	response, err := client.Do(request)
	if err != nil {
		log.Errorf("Get json response err: %v", err)
		return nil, err
	}
	defer response.Body.Close()

	respBody, _ := ioutil.ReadAll(response.Body)
	if response.StatusCode >= http.StatusOK && response.StatusCode <= http.StatusIMUsed {
		// limit the output
		outLen := len(respBody)
		truncated := 0
		if outLen > 200 {
			truncated = outLen - 200
			outLen = 200
		}
		log.Infof("Got response: url=%v, body=%v, tructated=%d", url, string(respBody[:outLen]), truncated)
		return respBody, nil
	}

	log.Errorf("Get fail response：code=%v, msg=%v", response.StatusCode, response.Status)

	if response.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("Unauthorized")
	}

	return nil, fmt.Errorf("%d(%s)", response.StatusCode, response.Status)
}
