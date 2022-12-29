package tr

import "time"

type scope struct {
	ctx  Ctx
	name string
	t0   time.Time
	w    int64
}

func NewScope(c Ctx, name string, w int64) *scope {
	return &scope{ctx: c, name: name, t0: time.Now(), w: w}
}

func (s *scope) Done() { s.ctx.Out(s.name, time.Since(s.t0), s.w) }
