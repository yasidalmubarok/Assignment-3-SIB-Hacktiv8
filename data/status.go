package data

import (
	"fmt"
	"math/rand"
)

type Status struct {
	WaterStatus int `json:"water"`
	WindStatus  int `json:"wind"`
}

var water, wind int

func UpdateData() Status {
	water = rand.Intn(100) + 1
	wind = rand.Intn(100) + 1

	return Status{
		WaterStatus: water,
		WindStatus: wind,
	}
}

func GetData() {
	waterStatus := "aman"

	if water < 6 {
		waterStatus = "aman"
	} else if water >= 5 && water <= 8 {
		waterStatus = "siaga"
	} else {
		waterStatus = "bahaya"
	}

	windStatus := "aman"
	if wind < 6 {
		windStatus = "aman"
	} else if wind >= 6 && wind <= 15 {
		windStatus = "siaga"
	} else {
		windStatus = "bahaya"
	}

	fmt.Println("status water: ", waterStatus)
	fmt.Println("status water: ", windStatus)
}
