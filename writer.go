package gel

import (
	"io"
)

// Writer is a type that slightly simplifies writing to an io.Writer
type Writer struct {
	io.Writer
}

// Write accepts variable number of bytes
func (w Writer) Write(s string) {
	_, err := w.Writer.Write([]byte(s))
	if err != nil {
		panic(err)
	}
}
