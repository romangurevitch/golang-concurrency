# 4. Finally, something works as expected

* [Simple counter](counter/simple.md)
* [Safe counter](counter/safe.md)

```go
package concurrency

// FinallySomethingWorksAsExpected but is it?
func FinallySomethingWorksAsExpected() int {
	c := &counter.SimpleCounter{}
	wg := sync.WaitGroup{}
	lock := sync.Mutex{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			lock.Lock()
			c.Inc()
			lock.Unlock()
		}()
	}

	wg.Wait()
	return c.Count()
}
```

```bash
 go test ../internal/concurrency -v -count=1 -run="FinallySomethingWorksAsExpected$" 
```

```bash
 go test ../internal/concurrency -v -count=1 -run="FinallySomethingWorksAsExpected$" -race 
```

## Result?

|                                                   Correctness                                                    |                                                   Consistency                                                    |                                                   Completeness                                                   |
|:----------------------------------------------------------------------------------------------------------------:|:----------------------------------------------------------------------------------------------------------------:|:----------------------------------------------------------------------------------------------------------------:|
| <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/question.svg" width="40"/> | <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/question.svg" width="40"/> | <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/question.svg" width="40"/> |

[Solution](example_4_solution.md)