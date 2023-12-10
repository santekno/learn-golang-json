package main

import (
	"encoding/json"
	"os"
)

func LogJSON(data interface{}) string {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

type Customer struct {
	FirstName  string   `json:"first_name"`
	MiddleName string   `json:"middle_name"`
	LastName   string   `json:"last_name"`
	Hobbies    []string `json:"hobbies"`
}

func GenerateObjectJSON(data Customer) string {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func ConvertObjectJSON(data string) Customer {
	var cust Customer
	err := json.Unmarshal([]byte(data), &cust)
	if err != nil {
		panic(err)
	}

	return cust
}

func ConvertMapJSON(data string) map[string]interface{} {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		panic(err)
	}

	return result
}

func DecodeStreamReaderJSON(file string) Customer {
	reader, _ := os.Open(file)
	decoder := json.NewDecoder(reader)

	var customer Customer
	err := decoder.Decode(&customer)
	if err != nil {
		panic(err)
	}
	return customer
}

func EncoderStreaWriterJSON(cust Customer) {
	writer, _ := os.Create("sample_output.json")
	encoder := json.NewEncoder(writer)

	err := encoder.Encode(cust)
	if err != nil {
		panic(err)
	}
}
