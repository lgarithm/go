package xterm

import (
	"bytes"
	"fmt"
)

type color struct {
	f uint8
	b uint8
}

// Standard XTerm Colors
var (
	Green = color{f: 32, b: 1}
	Red   = color{f: 35, b: 1}
)

func (c color) bs(text string) *bytes.Buffer {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "\x1b[%d;%dm", c.b, c.f)
	buf.WriteString(text)
	buf.WriteString("\x1b[m")
	return buf
}

func (c color) B(text string) []byte {
	return c.bs(text).Bytes()
}

func (c color) S(text string) string {
	return c.bs(text).String()
}
