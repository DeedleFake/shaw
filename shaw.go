package shaw

import (
	"bufio"
	"bytes"
	"io"
	"iter"
	"unsafe"

	"deedles.dev/shaw/dict"
)

// This might be one of the worst things that I've ever written.
type translator struct {
	r    *bufio.Reader
	dict *dict.Dict
	buf  bytes.Buffer
	err  error
}

// Translate returns an io.Reader that translates data coming from r
// as it itself is read from.
func Translate(r io.Reader, dict *dict.Dict) io.Reader {
	return &translator{
		r:    bufio.NewReader(r),
		dict: dict,
	}
}

func (t *translator) Read(buf []byte) (n int, err error) {
	for {
		bn, _ := t.buf.Read(buf[n:])
		n += bn
		if n == len(buf) {
			return n, nil
		}
		if t.err != nil {
			return n, t.err
		}

		t.advance()
	}
}

func (t *translator) advance() {
	t.buf.Reset()

	next, stop := iter.Pull(bytesseq(t.r, &t.err))
	defer stop()

	c, ok := next()
	if !ok {
		return
	}
	t.buf.WriteByte(c)

	translatable := isTranslatable(c)
	defer t.translate(translatable)

	for {
		c, ok := next()
		if !ok {
			return
		}
		if isTranslatable(c) != translatable {
			return
		}
		t.buf.WriteByte(c)
	}
}

func (t *translator) translate(translatable bool) {
	if !translatable {
		return
	}

	word := unsafe.String(unsafe.SliceData(t.buf.Bytes()), t.buf.Len())
	word = t.dict.Translate(word)
	t.buf.Reset()
	t.buf.WriteString(word)
}

// isTranslatable returns true if r is something expected in a
// translatable word.
func isTranslatable(r byte) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '\''
}

func bytesseq(r io.ByteScanner, perr *error) iter.Seq[byte] {
	return func(yield func(byte) bool) {
		for {
			c, err := r.ReadByte()
			if err != nil {
				*perr = err
				return
			}
			if !yield(c) {
				r.UnreadByte()
				return
			}
		}
	}
}
