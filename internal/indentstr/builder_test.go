package indentstr_test

import (
	"testing"

	"github.com/adzil/eqtest/internal/indentstr"
)

func TestBuilder_WriteString(t *testing.T) {
	var buf indentstr.Builder

	buf.WriteString("hello\nworld")
	buf.SetIndent(4)
	buf.WriteString("\nhello\nworld")
	buf.SetIndent(0)
	buf.WriteString("\nhello\nworld")

	if out := buf.String(); out != "hello\nworld\n    hello\n    world\nhello\nworld" {
		t.Error("unexpected indentation output:", out)
	}
}

func TestBuilder_Write(t *testing.T) {
	var buf indentstr.Builder

	// No need to test complicated scenarios as it is already tested by the
	// TestBuilder_WriteString function.
	buf.SetIndent(4)
	buf.Write([]byte("\nhello world"))

	if out := buf.String(); out != "\n    hello world" {
		t.Error("unexpected indentation output:", out)
	}
}

func TestBuilder_WriteByte(t *testing.T) {
	var buf indentstr.Builder

	buf.SetIndent(4)
	buf.WriteByte('\n')
	buf.WriteByte('a')

	if out := buf.String(); out != "\n    a" {
		t.Error("unexpected indentation output:", out)
	}
}

func TestBuilder_WriteRune(t *testing.T) {
	var buf indentstr.Builder

	buf.SetIndent(4)
	buf.WriteRune('\n')
	buf.WriteRune('a')

	if out := buf.String(); out != "\n    a" {
		t.Error("unexpected indentation output:", out)
	}
}
