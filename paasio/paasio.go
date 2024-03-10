// Package paasio contains various tools to report network IO statistics.
package paasio

import (
	"io"
	"sync"
)

// type readCounter encloses a io.Reader while also keeping track of all calls and reads to it.
type readCounter struct {
	reader io.Reader
	count  int64
	calls  int
	mx     sync.Mutex
}

// type writeCounter encloses a io.Writer while also keeping track of all calls and writes to it.
type writeCounter struct {
	writer io.Writer
	count  int64
	calls  int
	mx     sync.Mutex
}

// type readWriteCounter encompasses both readCounter and writeCounter types.
type readWriteCounter struct {
	ReadCounter
	WriteCounter
}

// NewWriteCounter creates a new WriteCounter.
func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{writer: writer}
}

// NewWriteCounter creates a new ReadCounter.
func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{reader: reader}
}

// NewReadWriteCounter creates a new ReadWriteCounter.
func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{NewReadCounter(readwriter), NewWriteCounter(readwriter)}
}

// Read method reads into the provided buffer.
func (rc *readCounter) Read(p []byte) (int, error) {
	rc.mx.Lock()
	defer rc.mx.Unlock()

	n, err := rc.reader.Read(p)

	rc.count += int64(n)
	rc.calls++

	return n, err
}

// ReadCount method returns the number of reads as well as number of calls to readCounter.
func (rc *readCounter) ReadCount() (int64, int) {
	rc.mx.Lock()
	defer rc.mx.Unlock()
	return rc.count, rc.calls
}

// Write method writes from the provided buffer.
func (wc *writeCounter) Write(p []byte) (int, error) {
	wc.mx.Lock()
	defer wc.mx.Unlock()

	n, err := wc.writer.Write(p)

	wc.count += int64(n)
	wc.calls++

	return n, err
}

// WriteCount method returns the number of writes as well as number of calls to writeCounter.
func (wc *writeCounter) WriteCount() (int64, int) {
	wc.mx.Lock()
	defer wc.mx.Unlock()
	return wc.count, wc.calls
}
