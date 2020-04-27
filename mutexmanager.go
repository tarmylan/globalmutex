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

func (p *MutexManager) lock(key interface{}) {
	p.gm.Lock(key)
}

func (p *MutexManager) unlock(key interface{}) {
	p.gm.Unlock(key)
}

func (p *MutexManager) rLock(key interface{}) {
	p.gm.RLock(key)
}

func (p *MutexManager) rUnlock(key interface{}) {
	p.gm.RUnlock(key)
}

func (p *MutexManager) execute(key interface{}, f func()) {
	p.gm.Lock(key)
	defer p.gm.Unlock(key)
	f()
}

func Lock(key interface{}) {
	_global_mutex_manager.lock(key)
}

func Unlock(key interface{}) {
	_global_mutex_manager.unlock(key)
}

func RLock(key interface{}) {
	_global_mutex_manager.rLock(key)
}

func RUnlock(key interface{}) {
	_global_mutex_manager.rUnlock(key)
}

func LockDo(key interface{}, f func()) {
	_global_mutex_manager.execute(key, f)
}
