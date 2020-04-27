package globalmutex

import (
	"sync"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestMain(m *testing.M) {
	var (
		wg sync.WaitGroup
	)

	count := 10000
	var lock_count int
	lockDo := func(key string) {
		go func() {
			wg.Add(1)
			LockDo(key, func() {
				lock_count++
			})
			wg.Done()
		}()
	}

	counter := func(key string) {
		for i := 0; i < count; i++ {
			lockDo(key)
		}
	}

	counter("a")
	wg.Wait()
	log.Infof("in %v out %v", count, lock_count)
}
