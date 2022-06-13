package counter

import "sync"

type ThreadSafe struct {
	counter int

	lock sync.RWMutex
}

func (s *ThreadSafe) Inc() {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.counter++
}

func (s *ThreadSafe) Count() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.counter
}
