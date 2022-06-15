# 5. Non-stopping go routines

[Safe counter](counter/safe.md)

```go
package concurrency

// WorkingEndlessly is that a good idea?
func WorkingEndlessly() int {
	c := &counter.SafeCounter{}
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		for {
			c.Inc()
		}
	}()

	wg.Wait()
	return c.Count()
}
```

```bash
 go test ../internal/concurrency -v -run="WorkingEndlessly$" 
```

```bash
 go test ../internal/concurrency -v -run="WorkingEndlessly$" -race 
```

## Result?

|                                                 Correctness                                                 |                                                 Consistency                                                 |                                                Completeness                                                 |
|:-----------------------------------------------------------------------------------------------------------:|:-----------------------------------------------------------------------------------------------------------:|:-----------------------------------------------------------------------------------------------------------:|
| <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/yes.png" width="40"/> | <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/yes.png" width="40"/> | <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/yes.png" width="40"/> |

[Next example](example_6.md)