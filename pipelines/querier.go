// +build OMIT

package main

import "encoding/json"

func main() {
}

type record struct {
}

// START OMIT
func query(key string) <-chan record {
	recordsC := make(chan record)

	go func() {
		defer close(recordsC)

		repo.Query(key, func(row map[string]interface{}) {
			r := record{}
			json.Unmarshal(row, &r)
			recordsC <- r
		})
	}()

	return recordsC
}

// STOP OMIT
