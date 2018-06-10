package control

import "sync"

// WithLock runs a func guarded by lock
func WithLock(l sync.Locker, f func()) {
	l.Lock()
	defer l.Unlock()
	f()
}
