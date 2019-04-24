package exercises

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (r MyReader) Read(buf []byte) (int, error) {
	if cap(buf) < 1 {
		return 0, io.ErrShortBuffer
	}

	buf[0] = 'A'

	return 1, nil
}

func myReaderTest() {
	reader.Validate(MyReader{})
}

type rot13Reader struct {
	r io.Reader
}

func rot13(char byte) byte {
	var rangeLowBound, rangeHighBound byte

	switch {
	case 'a' <= char && char <= 'z':
		rangeLowBound, rangeHighBound = 'a', 'z'
	case 'A' <= char && char <= 'Z':
		rangeLowBound, rangeHighBound = 'A', 'Z'
	default:
		return b
	}

	// x % y (modular) can be used as abstract solution
	// it determines how x is fits in [0, y] cyclic range
	// for example, 5 % 10 == 5, 5 is exists in range
	// 12 % 10 == 2 - can be treated as "we broke upper bound and starting again from zero"
	// if our low bound starts with >0 value, we need to adjust it temporarily
	// to 0 while calculation and apply it right after
	return (b-rangeLowBound+13)%(rangeHighBound-rangeLowBound+1) + rangeLowBound
}

func (r *rot13Reader) Read(outputBuffer []byte) (bytesRead int, err error) {
	bytesRead, err = r.r.Read(outputBuffer)

	for i := 0; i < bytesRead; i++ {
		outputBuffer[i] = rot13(outputBuffer[i])
	}

	return
}

func rot13ReaderTest() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

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

	myReaderTest()
	rot13ReaderTest()
}
