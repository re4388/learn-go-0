package learn

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func CreateFile(fileName string) {
	file, _ := os.Create(fileName)
	fmt.Fprint(file, "qqqq")
	file.Close()
}

func AppendFile(fn, text string) error {
	file, err := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening file for writing: %w", err)
	}

	defer file.Close()

	if _, err := file.Write([]byte(text)); err != nil {
		return fmt.Errorf("error writing text to fiel: %w", err)
	}

	return nil
}

/*

The io package specifies the io.Reader interface,
which represents the read end of a stream of data.

The Go standard library contains many implementations of this interface,
including files, network connections, compressors, ciphers, and others.


The io.Reader interface has a Read method:
func (T) Read(b []byte) (n int, err error)

Read populates the given byte slice with data and returns the number of bytes populated and an error value.
It returns an io.EOF error when the stream ends.

The example code creates a strings.Reader and consumes its output 8 bytes at a time.
*/

func Reader() {
	reader := strings.NewReader("Hello, Reader!")
	byte := make([]byte, 8)

	for {
		n, err := reader.Read(byte)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, byte)
		fmt.Printf("b[:n] = %q\n", byte[:n])

		if err == io.EOF {
			break
		}
	}
}
