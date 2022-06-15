# Let's make a small change :) 

```go
package counter

type SimpleCounter struct {
	Counter int
}
```

```go
package concurrency

// LetsMakeASmallChange ohh no!
func LetsMakeASmallChange() int {
	c := &counter.SimpleCounter{}
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
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
 go test ../internal/concurrency -v -run="LetsMakeASmallChange$" 
```

```bash
 go test ../internal/concurrency -v -run="LetsMakeASmallChange$" -race 
```

## Result?

|                                                   Correctness                                                    |                                                   Consistency                                                    |                                                   Completeness                                                   |
|:----------------------------------------------------------------------------------------------------------------:|:----------------------------------------------------------------------------------------------------------------:|:----------------------------------------------------------------------------------------------------------------:|
| <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/question.svg" width="40"/> | <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/question.svg" width="40"/> | <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/question.svg" width="40"/> |

[Solution](example_3_solution.md)