# 5. Non-stopping go routines

[Safe counter](counter/safe.md)

```go
package concurrency

// NonStoppingGoRoutine is that a good idea?
func NonStoppingGoRoutine() int {
	c := &counter.SafeCounter{}
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		for {
			inlinePrint(c.Inc())
		}
	}()

	wg.Wait()
	return c.Count()
}
```

```bash
 go test ../internal/concurrency -v -count=1 -run="NonStoppingGoRoutine$" 
```

```bash
 go test ../internal/concurrency -v -count=1 -run="NonStoppingGoRoutine$" -race 
```

## Result?

|                                                   Correctness                                                    |                                                   Consistency                                                    |                                                   Completeness                                                   |
|:----------------------------------------------------------------------------------------------------------------:|:----------------------------------------------------------------------------------------------------------------:|:----------------------------------------------------------------------------------------------------------------:|
| <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/question.svg" width="40"/> | <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/question.svg" width="40"/> | <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/question.svg" width="40"/> |

[Solution](example_5_solution.md)