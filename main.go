package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const REST_URL = "http://128.131.132.179:8080/lpims/rest/v1/"
const CALLURL_REACTORS = REST_URL + "reactors/"

var USERNAME string = "student"
var PASSWORD string = "bioVT"
var test string

type Reactor struct {
	Data []struct {
		Id      uint    `json:"id"`
		Name    string  `json:"name"`
		Volume  float32 `json:"volume"`
		Running bool    `json:"running"`
	} `json:"data"`
}

func main() {
	fmt.Println("Hello, go")
	fmt.Println(REST_URL)
	get_reactors()
}

func get_reactors() {
	var reactors Reactor
	body := get_request(CALLURL_REACTORS)
	err := json.Unmarshal(body, &reactors)
	if err != nil {
		fmt.Println("Error unmarshaling json:", err)
		return
	}
	for _, element := range reactors.Data {
		fmt.Println("ID:", element.Id, "Name:", element.Name, "Volume:", element.Volume, "Running:", element.Running)
	}
}

func get_request(callurl string) []byte {
	client := &http.Client{}

	req, err := http.NewRequest("GET", callurl, nil)
	if err != nil {
		fmt.Println("ERR!:", err)
	}
	req.SetBasicAuth(USERNAME, PASSWORD)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil
	}
	defer resp.Body.Close()

	// Check the HTTP status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("HTTP Error:", resp.Status)
		return nil
	}

	// Read the HTTP response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil
	}
	return body
	// Print the HTTP response body
	//fmt.Println("Response Body:", string(body))
	//fmt.Print(response)
}

// function get_reactor_ids()
// calling_str = rest_url * "reactors/"
// response = HTTP.request("GET", calling_str, headers = ["Authorization" => "Basic $(auth)"])
// jsondata = JSON.parse(String(response.body))["data"]
// name = [reactor["name"] for reactor in jsondata]
// id = [reactor["id"] for reactor in jsondata]
// return Dict(zip(name,id))
// end
