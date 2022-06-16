package counter

type SimpleCounter struct {
	counter int
}

func (s *SimpleCounter) Inc() int {
	s.counter++
	return s.counter
}

func (s *SimpleCounter) Count() int {
	return s.counter
}
