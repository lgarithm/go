package profile

import "io"

type Scope interface {
	Done()
}

type Profiler interface {
	Profile(string) Scope
	WriteSummary(io.Writer) error
}

func New() Profiler {
	return NewSimpleProfiler()
}
