# 1. Let's start with a simple example

[Simple counter](counter/simple.md)

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
 go test ../internal/concurrency -v -count=1 -run="UnexpectedResult$" 
```

```bash
 go test ../internal/concurrency -v -count=1 -run="UnexpectedResult$" -race 
```

## Result?

|                                                   Correctness                                                    |                                                   Consistency                                                    |                                                   Completeness                                                   |
|:----------------------------------------------------------------------------------------------------------------:|:----------------------------------------------------------------------------------------------------------------:|:----------------------------------------------------------------------------------------------------------------:|
| <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/question.svg" width="40"/> | <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/question.svg" width="40"/> | <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/question.svg" width="40"/> |

[Solution](example_1_solution.md)