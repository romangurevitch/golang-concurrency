# 6. Let's handle shutdown gracefully?

[Safe counter](counter/safe.md)

```go
package concurrency

// WorkingEndlesslyWithAWayOut is it good enough though?
func WorkingEndlesslyWithAWayOut() int {
	c := &counter.SafeCounter{}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			c.Inc()
		}
	}()

	fmt.Println("Waiting...")
	<-sigs

	return c.Count()
}
```

```bash
 go test ../internal/concurrency -v -run="WorkingEndlesslyWithAWayOut$" 
```

```bash
 go test ../internal/concurrency -v -run="WorkingEndlesslyWithAWayOut$" -race 
```

## Result?

|                                                 Correctness                                                 |                                                 Consistency                                                 |                                                Completeness                                                 |
|:-----------------------------------------------------------------------------------------------------------:|:-----------------------------------------------------------------------------------------------------------:|:-----------------------------------------------------------------------------------------------------------:|
| <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/yes.png" width="40"/> | <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/yes.png" width="40"/> | <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/yes.png" width="40"/> |

[Next example](example_7.md)