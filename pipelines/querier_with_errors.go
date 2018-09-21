// +build OMIT

package main

import "encoding/json"

func main() {
}

type record struct {
}

// START OMIT
func query(key string) (<-chan record, <-chan error) {
	recordsC := make(chan record, 10)
	errsC := make(chan error, 1)

	go func() {
		defer close(recordsC)

		repo.Query(key, func(row map[string]interface{}) {
			r := record{}
			err := json.Unmarshal(row, &r)

			if err != nil {
				errsC <- err
				return
			}

			recordsC <- r
		})
	}()

	return recordsC, errsC
}

// STOP OMIT
