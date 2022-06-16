package counter

import "sync"

type SafeCounter struct {
	counter int

	lock sync.RWMutex
}

func (s *SafeCounter) Inc() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.counter++
	return s.counter
}

func (s *SafeCounter) Count() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.counter
}
