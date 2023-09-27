package iox

import "io"

type ByteCountingWriter struct {
	n    int64
	base io.Writer
}

func (b *ByteCountingWriter) Write(buf []byte) (int, error) {
	n, err := b.base.Write(buf)
	b.n += int64(n)
	return n, err
}

func (b *ByteCountingWriter) NumBytes() int64 {
	return b.n
}

func NewByteCountingWriter(base io.Writer) *ByteCountingWriter {
	return &ByteCountingWriter{
		base: base,
		n:    0,
	}
}

type ByteCountingReader struct {
	n    int64
	base io.Reader
}

func (b *ByteCountingReader) Read(buf []byte) (int, error) {
	n, err := b.base.Read(buf)
	b.n += int64(n)
	return n, err
}

func (b *ByteCountingReader) NumBytes() int64 {
	return b.n
}

func NewByteCountingReader(base io.Reader) *ByteCountingReader {
	return &ByteCountingReader{
		base: base,
		n:    0,
	}
}
