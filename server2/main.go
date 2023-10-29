package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type Data struct {
	Data1 string `json:"data1"`
	Data2 string `json:"data2"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/data", dataHandler)

	http.ListenAndServe(":85", r)

}

func dataHandler(responseWriter http.ResponseWriter, request *http.Request) {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		data, err := readDataFromJson()
		jsonData, err := json.Marshal(data)
		if err != nil {
			http.Error(responseWriter, "JSON marshal error", http.StatusInternalServerError)
			return
		}
		if err != nil {

		}

		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.Write(jsonData)
		//time.Sleep(2 * time.Second)
	}

}

func readDataFromJson() (Data, error) {
	// Read the JSON file
	jsonData, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return Data{}, err
	}

	// Create a variable of the Data struct
	var data Data

	// Unmarshal JSON data into the struct
	if err := json.Unmarshal(jsonData, &data); err != nil {
		fmt.Println("JSON unmarshal error:", err)
		return Data{}, err
	}

	// Access the data in the struct
	fmt.Printf("Name: %s, Value: %s\n", data.Data1, data.Data2)

	return data, nil

}
