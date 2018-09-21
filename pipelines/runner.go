// +build OMIT

package main

func main() {
}

// START OMIT
func runner(key string) error {
	records, _ := query(key)

	reader, _ = process(records)

	upload(reader)

	// lol what errors ¯\_(ツ)_/¯
}

// STOP OMIT
