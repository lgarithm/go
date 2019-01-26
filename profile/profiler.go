package profile

import "io"

type Scope interface {
	WithAmount(int64) Scope
	Done()
}

type Profiler interface {
	Profile(string) Scope
	WriteSummary(io.Writer) error
}

func New() Profiler {
	return NewSimpleProfiler()
}
