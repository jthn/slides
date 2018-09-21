// +build OMIT

package main

import "log"

func main() {
}

// START OMIT
func runner(key string) error {
	records, recordErrs := query(key)

	done, errs := process(records)

	upload(reader)

	select {
	case err := <-recordErrs:
		log.Fatal(err)
	case err := <-errs:
		log.Fatal(err)
	case <-done:
		return
	}
}

// STOP OMIT
