package data

import (
	"encoding/json"
	"fmt"
	"log"
)

func PrintStatus() {
	status := UpdateData()
	jsonStatus, err := json.MarshalIndent(status, "", "  ")

	if err != nil {
		log.Panicf("Error while printing status:%s\n", err.Error())
	}

	fmt.Println(string(jsonStatus))
}
