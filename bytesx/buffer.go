package bytesx

import (
	"bytes"
	"errors"
	"io"
)

type Buffer struct {
	buf []byte
	i   int64
}

func (ws *Buffer) Read(b []byte) (int, error) {
	if ws.i >= int64(len(ws.buf)) {
		return 0, io.EOF
	}
	n := copy(b, ws.buf[ws.i:])
	ws.i += int64(n)
	return n, nil
}

func (ws *Buffer) Write(b []byte) (int, error) {
	n := len(b)
	diff := ws.i - int64(len(ws.buf))

	var tail []byte
	if n+int(ws.i) < len(ws.buf) {
		tail = ws.buf[n+int(ws.i):]
	}

	if diff > 0 {
		ws.buf = append(ws.buf, append(bytes.Repeat([]byte{0o0}, int(diff)), b...)...)
		ws.buf = append(ws.buf, tail...)
	} else {
		ws.buf = append(ws.buf[:ws.i], b...)
		ws.buf = append(ws.buf, tail...)
	}

	ws.i += int64(n)

	return n, nil
}

// Copied from Go stdlib bytes.Reader
func (ws *Buffer) Seek(offset int64, whence int) (int64, error) {
	var abs int64
	switch whence {
	case 0:
		abs = offset
	case 1:
		abs = int64(ws.i) + offset
	case 2:
		abs = int64(len(ws.buf)) + offset
	default:
		return 0, errors.New("bytes.Reader.Seek: invalid whence")
	}
	if abs < 0 {
		return 0, errors.New("bytes.Reader.Seek: negative position")
	}
	ws.i = abs
	return abs, nil
}

func (ws *Buffer) Close() error {
	return nil
}

func NewBuffer(buf []byte) *Buffer {
	return &Buffer{
		buf: buf,
		i:   0,
	}
}
