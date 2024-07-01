package indentstr

import (
	"strings"
	"unsafe"
)

type builder = strings.Builder

type Builder struct {
	builder
	indent int
}

func (b *Builder) writeIndent() {
	for range b.indent {
		b.builder.WriteByte(' ')
	}
}

func (b *Builder) SetIndent(n int) {
	b.indent = n
}

func (b *Builder) WriteString(s string) (int, error) {
	n := len(s)

	for len(s) > 0 {
		var noNewline bool
		pos := strings.IndexByte(s, '\n')
		if pos < 0 {
			pos = len(s)
			noNewline = true
		} else {
			pos += 1
		}

		b.builder.WriteString(s[:pos])
		s = s[pos:]

		if !noNewline {
			b.writeIndent()
		}
	}

	return n, nil
}

func (b *Builder) WriteByte(v byte) error {
	b.builder.WriteByte(v)

	if v == '\n' {
		b.writeIndent()
	}

	return nil
}

func (b *Builder) WriteRune(r rune) (int, error) {
	n, _ := b.builder.WriteRune(r)

	if r == '\n' {
		b.writeIndent()
	}

	return n, nil
}

func (b *Builder) Write(v []byte) (n int, err error) {
	s := unsafe.String(unsafe.SliceData(v), len(v))
	return b.WriteString(s)
}
