package main

import "fmt"

// DataStream structure represents a data stream
type DataStream struct {
	data []map[string]int
}

// NewDataStream is a constructor for DataStream
func NewDataStream(data []map[string]int) *DataStream {
	return &DataStream{data: data}
}

// TODO: Implement a method 'SliceMiddle' that calculates and returns the middle element of a slice (given `start` and `end` positions)
func (ds *DataStream) SliceMiddle(start, end int) map[string]int {
	MiddleElement := (start + end) / 2
	return ds.data[MiddleElement]
}

func main() {
	marketIndices := []map[string]int{
		{"id": 0, "value": 33000},
		{"id": 1, "value": 14000},
		{"id": 2, "value": 4200},
		{"id": 3, "value": 2300},
	}
	stream := NewDataStream(marketIndices)

	// TODO: Retrieve and log the name of the middle element of a slice that comprises the first 3 elements in the stream
	fmt.Println(stream.SliceMiddle(0, 3))
}
