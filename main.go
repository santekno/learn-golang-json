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
