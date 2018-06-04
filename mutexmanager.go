package globalmutex

import (
	"sync"
)

var (
	once                  sync.Once
	_global_mutex_manager MutexManager
)

type MutexManager struct {
	gm GlobalMutex
}

func init() {
	once.Do(func() {
		_global_mutex_manager.init()
	})
}

func (p *MutexManager) init() {
	p.gm = NewGlobalMutex()
}

func (p *MutexManager) lock(key string) {
	p.gm.Lock(key)
}

func (p *MutexManager) unlock(key string) {
	p.gm.Unlock(key)
}

func (p *MutexManager) rLock(key string) {
	p.gm.RLock(key)
}

func (p *MutexManager) rUnlock(key string) {
	p.gm.RUnlock()
}

func (p *MutexManager) execute(key string, f func()) {
	p.gm.Lock(key)
	defer p.gm.Unlock(key)
	f()
}

func Lock(key string) {
	_global_mutex_manager.lock(key)
}

func Unlock(key string) {
	_global_mutex_manager.unlock(key)
}

func RLock(key string) {
	_global_mutex_manager.rLock(key)
}

func RUnlock(key string) {
	_global_mutex_manager.rUnlock(key)
}

func LockDo(key string, f func()) {
	_global_mutex_manager.execute(key, f)
}
