package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GETJson(url string, data interface{}) bool {
	req, err := http.NewRequest("GET", url, nil)
	// req.Header.Add("Authorization", "Basic "+basicAuth(u, p))
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		print("GET error: ", err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("%s: did not get acceptable status code: %v body: %q",
			url, resp.Status, string(body))
		return false
	}
	err = json.NewDecoder(resp.Body).Decode(data)
	if err != nil {
		print("decode error: ", err.Error())
		return false
	}

	return true
}
