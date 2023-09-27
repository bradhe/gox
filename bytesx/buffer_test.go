package bytesx

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuffer(t *testing.T) {
	t.Run("can write and read from buffer", func(t *testing.T) {
		buf := NewBuffer([]byte{})

		t.Run("can write to buffer", func(t *testing.T) {
			n, err := buf.Write([]byte("Hello, world!"))
			assert.NoError(t, err)
			assert.Equal(t, n, 13)
		})

		t.Run("can rewind buffer", func(t *testing.T) {
			n, err := buf.Seek(0, io.SeekStart)
			assert.NoError(t, err)
			assert.Equal(t, n, int64(0))
		})

		t.Run("can read from buffer", func(t *testing.T) {
			var b [5]byte
			n, err := buf.Read(b[:])
			assert.NoError(t, err)
			assert.Equal(t, n, 5)
			assert.Equal(t, string(b[:]), "Hello")
		})
	})

	t.Run("can read from a seeded buffer", func(t *testing.T) {
		buf := NewBuffer([]byte("Hello, world!"))

		var b [5]byte
		n, err := buf.Read(b[:])
		assert.NoError(t, err)
		assert.Equal(t, n, 5)
		assert.Equal(t, string(b[:]), "Hello")
	})

	t.Run("allows content to be overwritten", func(t *testing.T) {
		buf := NewBuffer([]byte("Hello, world!"))

		n, err := buf.Seek(0, io.SeekStart)
		assert.NoError(t, err)
		assert.Equal(t, n, int64(0))

		_, err = buf.Write([]byte("Goodbye, my friend!"))
		assert.NoError(t, err)

		n, err = buf.Seek(0, io.SeekStart)
		assert.NoError(t, err)
		assert.Equal(t, n, int64(0))

		var b [7]byte
		_, err = buf.Read(b[:])
		assert.NoError(t, err)
		assert.Equal(t, string(b[:]), "Goodbye")
	})
}
