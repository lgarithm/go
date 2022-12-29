package tr

import (
	"log"
	"strings"
	"time"
)

type logCtx struct {
	Ctx
	depth int
}

func Log(ctx Ctx) Ctx {
	return &logCtx{Ctx: ctx}
}

func tab(n int) string { return strings.Repeat(`  `, n) }

func (c *logCtx) In(name string) {
	c.Ctx.In(name)
	log.Printf("%s{ // %s", tab(c.depth), name)
	c.depth++
}

func (c *logCtx) Out(name string, d time.Duration, w int64) {
	c.depth--
	log.Printf("%s} // %s", tab(c.depth), name)
	c.Ctx.Out(name, d, w)
}

func (c *logCtx) Trace(name string) Scope {
	s := NewScope(c, name, 0)
	c.In(name)
	return s
}

func (c *logCtx) TraceW(name string, w int64) Scope {
	s := NewScope(c, name, w)
	c.In(name)
	return s
}
