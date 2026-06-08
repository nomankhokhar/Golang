package main

import (
	"fmt"
)

type DataStream struct {
	data []map[string]int
}

func NewDataStream(data []map[string]int) *DataStream {
	return &DataStream{data: data}
}

func (ds *DataStream) SliceToString(start, end int) string {
	SliceOfStream := ""
	for i := start; i < end; i++ {
		SliceOfStream += fmt.Sprintf("{id: %d, value: %d }", ds.data[i]["id"], ds.data[i]["value"])

		if i < end-1 {
			SliceOfStream += ","
		}
	}
	return SliceOfStream
}

func main() {
	data := []map[string]int{
		{"id": 1, "value": 65},
		{"id": 2, "value": 75},
		{"id": 3, "value": 85},
		{"id": 4, "value": 95},
	}

	sensorReadings := NewDataStream(data)
	fmt.Println(sensorReadings.SliceToString(1, 3)) // Should print: {'id': 2, 'value': 75}, {'id': 3, 'value': 85}
}
