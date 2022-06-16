# 4. Finally, something works as expected - using safe counter

* [Simple counter](counter/simple.md)
* [Safe counter](counter/safe.md)

```go
package concurrency

// FinallySomethingWorksAsExpectedSafeCounter but is it?
func FinallySomethingWorksAsExpectedSafeCounter(count int) int {
	c := &counter.SafeCounter{}
	wg := sync.WaitGroup{}

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}

	wg.Wait()
	return c.Count()
}
```

```bash
 go test ../internal/concurrency -v -count=1 -run="FinallySomethingWorksAsExpectedSafeCounter$" 
```

```bash
 go test ../internal/concurrency -v -count=1 -run="FinallySomethingWorksAsExpectedSafeCounter$" -race 
```

## Result?

|                                                 Correctness                                                 |                                                 Consistency                                                 |                                                Completeness                                                 |
|:-----------------------------------------------------------------------------------------------------------:|:-----------------------------------------------------------------------------------------------------------:|:-----------------------------------------------------------------------------------------------------------:|
| <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/yes.png" width="40"/> | <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/yes.png" width="40"/> | <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/yes.png" width="40"/> |

[Next example](example_5.md)