# Let's start with a simple example

```go
package counter

type SimpleCounter struct {
	Counter int
}
```

```go
package concurrency

// UnexpectedResult what did you expect?
func UnexpectedResult() int {
	c := &counter.Simple{}

	go func() {
		for i := 0; i < 1000; i++ {
			c.Inc()
		}
	}()

	return c.Count()
}
```

```bash
 go test ../internal/concurrency -v -run="UnexpectedResult$" 
```

```bash
 go test ../internal/concurrency -v -run="UnexpectedResult$" -race 
```

## Result?

|                                                Correctness                                                 |                                                Consistency                                                 |                                                Completeness                                                |
|:----------------------------------------------------------------------------------------------------------:|:----------------------------------------------------------------------------------------------------------:|:----------------------------------------------------------------------------------------------------------:|
| <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/no.png" width="40"/> | <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/no.png" width="40"/> | <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/no.png" width="40"/> |

[Next example](example_2.md)