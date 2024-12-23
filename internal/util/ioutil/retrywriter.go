package ioutil

import (
	"io"
	"time"
)

type RetryWriter struct {
	io.Writer
}

func NewRetryWriter(writer io.Writer) *RetryWriter {
	return &RetryWriter{writer}
}

func (w *RetryWriter) Write(p []byte) (int, error) {
	ntotal := 0
	var err error
	var n int
	for len(p) > 0 {
		n, err = w.Writer.Write(p)
		ntotal += n
		p = p[n:]
		if n == 0 && err != io.ErrShortWrite {
			return ntotal, err
		}
		// to avoid busy waiting loops that use a lot of CPU
		time.Sleep(10 * time.Millisecond)
	}
	return ntotal, err
}
