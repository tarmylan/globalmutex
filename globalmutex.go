package globalmutex

import (
	"sync"
	"sync/atomic"
)

type GlobalMutex interface {
	Lock(interface{})
	RLock(interface{})
	Unlock(interface{})
	RUnlock(interface{})
}

type mutex struct {
	ref int64
	mu  *sync.RWMutex
}

type locker struct {
	inUse sync.Map
	pool  sync.Pool
}

func (l *locker) Lock(key interface{}) {
	m := l.get(key)
	atomic.AddInt64(&m.ref, 1)
	m.mu.Lock()
}

func (l *locker) RLock(key interface{}) {
	m := l.get(key)
	atomic.AddInt64(&m.ref, 1)
	m.mu.RLock()
}

func (l *locker) Unlock(key interface{}) {
	m := l.get(key)
	m.mu.Unlock()
	l.put(key, m)
}

func (l *locker) RUnlock(key interface{}) {
	m := l.get(key)
	m.mu.RUnlock()
	l.put(key, m)
}

func (l *locker) put(key interface{}, m *mutex) {
	atomic.AddInt64(&m.ref, -1)
	if m.ref <= 0 {
		l.pool.Put(m.mu)
		l.inUse.Delete(key)
	}
}

func (l *locker) get(key interface{}) *mutex {
	res, _ := l.inUse.LoadOrStore(key, &mutex{
		ref:    0,
		locker: l.pool.Get().(*sync.RWMutex),
	})

	return res.(*mutex)
}

func NewGlobalMutex() GlobalMutex {
	return &locker{
		pool: sync.Pool{
			New: func() interface{} {
				return &sync.RWMutex{}
			},
		},
	}
}
