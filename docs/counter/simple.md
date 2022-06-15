# Simple counter

```go
package counter

type SimpleCounter struct {
	counter int
}

func (s *SimpleCounter) Inc() {
	s.counter++
}

func (s *SimpleCounter) Count() int {
	return s.counter
}
```
