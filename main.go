package main

import (
	"assignment-3/data"
	"time"
)

func main() {
	go func() {
		for {
			data.UpdateData()
			data.PrintStatus()
			data.GetData()
			time.Sleep(15 * time.Second)
		}
	}()
	select{}
}