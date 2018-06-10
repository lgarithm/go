package sema

type lock chan struct{}

// A Sema represents an object that can be Acquire and Released
type Sema interface {
	Acquire() bool
	Release()
}

// New creates a Sema
func New() Sema {
	ch := make(lock, 1)
	ch <- struct{}{}
	return ch
}

func (l lock) Acquire() bool {
	select {
	case <-l:
		return true
	default:
		return false
	}
}

func (l lock) Release() {
	l <- struct{}{}
}
