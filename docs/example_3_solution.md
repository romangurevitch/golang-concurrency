# 3. Let's make a small change :)

[Simple counter](counter/simple.md)

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
 go test ../internal/concurrency -v -count=1 -run="LetsMakeASmallChange$" 
```

```bash
 go test ../internal/concurrency -v -count=1 -run="LetsMakeASmallChange$" -race 
```

## Result?

|                                                 Correctness                                                 |                                                 Consistency                                                 |                                                Completeness                                                 |
|:-----------------------------------------------------------------------------------------------------------:|:-----------------------------------------------------------------------------------------------------------:|:-----------------------------------------------------------------------------------------------------------:|
| <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/yes.png" width="40"/> | <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/yes.png" width="40"/> | <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/yes.png" width="40"/> |

[Next example](example_4.md)