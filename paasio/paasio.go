package paasio

import (
	"io"
	"sync"
)

// NewReadCounter creates a wrapper around io.Reader that calculates amount of bytes
// read and number of read operations
func NewReadCounter(r io.Reader) ReadCounter {
	return &readCounter{
		r: r,
	}
}

// NewWriteCounter creates a wrapper around io.Writer that calculates amount of bytes
// written and number of write operations
func NewWriteCounter(w io.Writer) WriteCounter {
	return &writeCounter{
		w: w,
	}
}

// NewReadWriteCounter creates a wrapper around io.ReadWriter that calculates number of bytes
// written/read and number of read/write operations
func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{
		readCounter{r: rw},
		writeCounter{w: rw},
	}
}

func (wc *writeCounter) Write(p []byte) (n int, err error) {

	n, err = wc.w.Write(p)

	if err == nil {

		wc.Lock()
		defer wc.Unlock()

		wc.ops++
		wc.bytes += int64(n)
	}

	return
}

func (wc *writeCounter) WriteCount() (n int64, nops int) {

	wc.Lock()
	defer wc.Unlock()

	return wc.bytes, int(wc.ops)
}

func (rc *readCounter) Read(p []byte) (n int, err error) {

	n, err = rc.r.Read(p)

	if err == nil {

		rc.Lock()
		defer rc.Unlock()

		rc.ops++
		rc.bytes += int64(n)
	}

	return
}

func (rc *readCounter) ReadCount() (n int64, nops int) {

	rc.Lock()
	defer rc.Unlock()

	return rc.bytes, int(rc.ops)
}

type counter struct {
	bytes int64
	ops   int64
	sync.Mutex
}

type writeCounter struct {
	counter
	w io.Writer
}

type readCounter struct {
	counter
	r io.Reader
}

type readWriteCounter struct {
	readCounter
	writeCounter
}
