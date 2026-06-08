package main

import (
	"fmt"
)

type DataStream struct {
	data []map[string]string
}

// TODO: Update this method to retrieve and return the last element of the data stream
func (ds DataStream) GetLast() map[string]string {
	fmt.Println(ds.data)
	return ds.data[len(ds.data)-1]
}

func main() {
	weatherData := []map[string]string{
		{"time": "10:00AM", "temperature": "15", "wind": "NE"},
		{"time": "11:00AM", "temperature": "17", "wind": "E"},
		{"time": "12:00PM", "temperature": "19", "wind": "SE"},
	}

	stream := DataStream{data: weatherData}

	firstData := stream.GetLast()
	fmt.Printf("time: %s, temperature: %s, wind: %s\n", firstData["time"], firstData["temperature"], firstData["wind"])
}
