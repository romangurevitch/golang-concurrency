# 2. Let's try and fix the issues

[Simple counter](counter/simple.md)

```go
package concurrency

// UnexpectedResultFix is it fixed?
func UnexpectedResultFix() int {
	c := &counter.SimpleCounter{}
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			c.Inc()
		}
	}()

	wg.Wait()
	return c.Count()
}
```

```bash
 go test ../internal/concurrency -v -run="UnexpectedResultFix$" 
```

```bash
 go test ../internal/concurrency -v -run="UnexpectedResultFix$" -race 
```

## Result?

|                                                 Correctness                                                 |                                                 Consistency                                                 |                                                Completeness                                                 |
|:-----------------------------------------------------------------------------------------------------------:|:-----------------------------------------------------------------------------------------------------------:|:-----------------------------------------------------------------------------------------------------------:|
| <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/yes.png" width="40"/> | <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/yes.png" width="40"/> | <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/yes.png" width="40"/> |

[Next example](example_3.md)