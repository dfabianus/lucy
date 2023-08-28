package lucrest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const REST_URL = "http://128.131.132.179:8080/lpims/rest/v1/"
const CALLURL_REACTORS = REST_URL + "reactors/"
const CALLURL_REACTORS_RUNNING = REST_URL + "reactors?running=true"
const CALLURL_PROC_RUNNING = REST_URL + "processes?running=true"
const CALLURL_PROC_NAME = REST_URL + "processes?name="
const CALLURL_PROC_ID = REST_URL + "processes?Id=" //
const CALLURL_PORT_NAME = REST_URL + "ports?name="
const CALLURL_SIGNALS_FROM_PROC_ID = REST_URL + "signals?processId="
const CALLURL_SIGNALS_FROM_SIGNAL_ID = REST_URL + "signals/"

//const CALLURL_SIGNALS_FROM_PROC_ID = REST_URL + "signals/"

var USERNAME string = "student"
var PASSWORD string = "bioVT"
var test string

type Reactor struct {
	Data []struct {
		Id      uint    `json:"id"`
		Name    string  `json:"name"`
		Volume  float32 `json:"volume"`
		Running bool    `json:"running"`
		// Process Process `json:"process"`
	} `json:"data"`
}
type Process struct {
	Data []struct {
		Id   uint   `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
}

type Signal struct {
	Data []struct {
		Id   uint `json:"id"`
		Port struct {
			Id   uint   `json:"id"`
			Name string `json:"name"`
		} `json:"Port"`
		ProcessId uint `json:"processId"`
		Reactor   struct {
			Id   uint   `json:"id"`
			Name string `json:"name"`
		} `json:"reactor"`
		Device struct {
			Id   uint   `json:"id"`
			Name string `json:"name"`
		} `json:"device"`
		SubDevice struct {
			Id   uint   `json:"id"`
			Name string `json:"name"`
		} `json:"subDevice"`
		ModifiedTimestamp string `json:"modifiedTimestamp"`
	} `json:"data"`
}

type singleSignal struct {
	Data struct {
		Id   uint `json:"id"`
		Port struct {
			Id   uint   `json:"id"`
			Name string `json:"name"`
		} `json:"Port"`
		RealTime bool        `json:"realTime"`
		Values   [][]float64 `json:"values"`
		Reactor  struct {
			Id   uint   `json:"id"`
			Name string `json:"name"`
		} `json:"reactor"`
		Device struct {
			Id   uint   `json:"id"`
			Name string `json:"name"`
		} `json:"device"`
		SubDevice struct {
			Id   uint   `json:"id"`
			Name string `json:"name"`
		} `json:"subDevice"`
		ModifiedTimestamp string `json:"modifiedTimestamp"`
	} `json:"data"`
}

// func main_arc() {
// 	fmt.Println("Hello, go")
// 	fmt.Println(REST_URL)
// 	get_process("Fermenter 22 2023 Week 10 006")
// }

func GetReactors() {
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

func GetRunningReactors() {
	var reactors Reactor
	body := get_request(CALLURL_REACTORS_RUNNING)
	err := json.Unmarshal(body, &reactors)
	if err != nil {
		fmt.Println("Error unmarshaling json:", err)
		return
	}
	for _, element := range reactors.Data {
		fmt.Println("ID:", element.Id, "Name:", element.Name, "Volume:", element.Volume, "Running:", element.Running)
	}
	//fmt.Println("Response Body:", string(body))
}

func GetRunningProcesses() {
	var processes Process
	body := get_request(CALLURL_PROC_RUNNING)
	err := json.Unmarshal(body, &processes)
	if err != nil {
		fmt.Println("Error unmarshaling json:", err)
		return
	}
	for _, element := range processes.Data {
		fmt.Println("ID:", element.Id, "Name:", element.Name)
	}
	//fmt.Println("Response Body:", string(body))
}

func GetProcessByName(process_name string) {
	var processes Process
	body := get_request(CALLURL_PROC_NAME + process_name)
	err := json.Unmarshal(body, &processes)
	if err != nil {
		fmt.Println("Error unmarshaling json:", err)
		return
	}
	for _, element := range processes.Data {
		fmt.Println("ID:", element.Id, "Name:", element.Name)
	}
	// fmt.Println("Response Body:", string(body))
}

func GetProcessById(process_id string) { // does not work yet
	var processes Process
	body := get_request(CALLURL_PROC_ID + process_id)
	err := json.Unmarshal(body, &processes)
	if err != nil {
		fmt.Println("Error unmarshaling json:", err)
		return
	}
	for _, element := range processes.Data {
		fmt.Println("ID:", element.Id, "Name:", element.Name)
	}
	//fmt.Println("Response Body:", string(body))
}

func GetSignalsByProcessId(process_id string) { // does not work yet
	var signals Signal
	body := get_request(CALLURL_SIGNALS_FROM_PROC_ID + process_id)
	fmt.Println(CALLURL_SIGNALS_FROM_PROC_ID + process_id)
	//body := get_request("http://128.131.132.179:8080/lpims/rest/v1/signals?processId=8983")
	//body := get_request(CALLURL_SIGNALS_FROM_PROC_ID + "12212")
	err := json.Unmarshal(body, &signals)
	if err != nil {
		fmt.Println("Error unmarshaling json:", err)
		return
	}
	for _, element := range signals.Data {
		fmt.Println("Signal-ID:", element.Id, "Port-ID:", element.Port.Id, "Name:", element.Port.Name, "Device-Name:", element.Device.Name, "Subdevice-Name:", element.SubDevice.Name)
	}
}

func GetSignalsBySignalId(signal_id string) { // does not work yet
	var signals singleSignal
	body := get_request(CALLURL_SIGNALS_FROM_SIGNAL_ID + signal_id)
	//body := get_request("http://128.131.132.179:8080/lpims/rest/v1/signals?processId=8983")
	//body := get_request(CALLURL_SIGNALS_FROM_PROC_ID + "12212")
	err := json.Unmarshal(body, &signals)
	if err != nil {
		fmt.Println("Error unmarshaling json:", err)
		return
	}
	element := signals.Data
	fmt.Println("Signal-ID:", element.Id, "Port-ID:", element.Port.Id, "Name:", element.Port.Name, "Device-Name:", element.Device.Name, "Subdevice-Name:", element.SubDevice.Name)
	// fmt.Println("Response Body:", string(body))

}

func get_request(callurl string) []byte {

	client := &http.Client{}

	req, err := http.NewRequest("GET", callurl, nil)
	if err != nil {
		fmt.Println("ERR!:", err)
	}
	req.SetBasicAuth(USERNAME, PASSWORD)
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	//req.Header.Add("Authorization", "Basic "+basicAuth(USERNAME, PASSWORD))
	//req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("Transfer-Encoding", "chunked")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil
	}
	defer resp.Body.Close()

	//Check the HTTP status code
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
