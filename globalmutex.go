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
	mu    sync.RWMutex
}

func (l *locker) Lock(key interface{}) {
	m := l.get(key, 1)
	m.mu.Lock()
}

func (l *locker) RLock(key interface{}) {
	m := l.get(key, 1)
	m.mu.RLock()
}

func (l *locker) Unlock(key interface{}) {
	m := l.get(key, 0)
	m.mu.Unlock()
	l.put(key, m)
}

func (l *locker) RUnlock(key interface{}) {
	m := l.get(key, 0)
	m.mu.RUnlock()
	l.put(key, m)
}

func (l *locker) put(key interface{}, m *mutex) {
	l.mu.Lock()
	defer l.mu.Unlock()

	atomic.AddInt64(&m.ref, -1)
	if m.ref <= 0 {
		l.pool.Put(m.mu)
		l.inUse.Delete(key)
	}
}

func (l *locker) get(key interface{}, delta int64) *mutex {
	l.mu.Lock()
	defer l.mu.Unlock()

	res, _ := l.inUse.LoadOrStore(key, &mutex{
		ref: 0,
		mu:  l.pool.Get().(*sync.RWMutex),
	})

	m := res.(*mutex)
	if delta != 0 {
		atomic.AddInt64(&m.ref, delta)
	}

	return m
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
