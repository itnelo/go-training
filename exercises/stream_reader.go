package exercises

import (
	"fmt"
	"io"
	"strings"
)

func streamReader() {
	var data string = "This is data to read."
	var reader io.Reader = strings.NewReader(data)

	var buffer = make([]byte, 4)

	for {
		bytesRead, err := reader.Read(buffer)

		// From io.Reader docs:
		// Callers should always process the n > 0 bytes returned before
		// considering the error err.
		if bytesRead > 0 {
			// Process received data batch.
			fmt.Printf("Reading %d bytes from string: %s\n", bytesRead, buffer[:bytesRead])
		}

		if err == io.EOF {
			break
		}
	}
}
