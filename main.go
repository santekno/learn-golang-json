package main

import (
	"encoding/json"
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
