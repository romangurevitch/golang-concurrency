# Simple counter

```go
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
```

<img height="360" src="https://media.giphy.com/media/APqEbxBsVlkWSuFpth/giphy.gif" width="389" alt="?"/>