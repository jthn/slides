// +build OMIT

package main

import "io"

func main() {
}

// START OMIT
func process(<-chan record) (io.Reader, <-chan error) {
	reader, writer := io.Pipe()
	errsC := make(chan error, 1)

	go func() {
		defer writer.Close()

		for r := range records {
			b, err := r.MarshalBinary()

			if err != nil {
				errsC <- err
				return
			}

			if _, err := writer.Write(b); err != nil {
				errsC <- err
				return
			}
		}
	}()

	return reader, errsC
}

// STOP OMIT
