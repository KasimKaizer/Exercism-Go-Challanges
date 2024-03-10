// Package circular contains various tools to work with circular buffer.
package circular

import "errors"

// type Buffer represents a circular buffer in memory.
type Buffer struct {
	Buf    []byte
	WIdx   int
	RIdx   int
	IsFull bool
}

var (
	errBufFull  = errors.New("circular.Buffer: buffer is full")
	errBufEmpty = errors.New("circular.Buffer: buffer is empty")
)

// NewBuffer creates a new circular buffer with the given length.
func NewBuffer(size int) *Buffer {
	newBuf := make([]byte, size)
	return &Buffer{Buf: newBuf}
}

// ReadByte method returns a single byte from the front of the buffer, buffer follows FIFO.
func (b *Buffer) ReadByte() (byte, error) {
	if !b.IsFull && b.RIdx == b.WIdx {
		return 0, errBufEmpty
	}

	out := b.Buf[b.RIdx]
	b.RIdx = (b.RIdx + 1) % len(b.Buf)
	b.IsFull = false
	return out, nil
}

// WriteByte method writes a byte to the end of the buffer.
func (b *Buffer) WriteByte(c byte) error {
	if b.IsFull {
		return errBufFull
	}
	b.Buf[b.WIdx] = c
	b.WIdx = (b.WIdx + 1) % len(b.Buf)
	b.IsFull = b.RIdx == b.WIdx
	return nil
}

// Overwrite method writes a byte to the end of the buffer, if the buffer is full then it would
// overwrite the oldest byte.
func (b *Buffer) Overwrite(c byte) {
	if b.IsFull {
		b.ReadByte()
	}
	b.WriteByte(c)
}

// Reset method resets the buffer.
func (b *Buffer) Reset() {
	b.IsFull = false
	b.RIdx, b.WIdx = 0, 0
}
