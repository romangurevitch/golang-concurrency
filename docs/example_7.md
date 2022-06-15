# # Let's handle shutdown gracefully, for real this time!

[Safe counter](counter/safe.md)

```go
package concurrency

// WorkingEndlesslyWithAGoodWayOut yes?
func WorkingEndlesslyWithAGoodWayOut() int {
	wg := sync.WaitGroup{}
	c := &counter.SafeCounter{}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-sigs:
				fmt.Println("\nSignal received, shutting down!")
				return
			default:
				c.Inc()
				inlinePrint(c.Count())
			}

		}
	}()

	fmt.Println("Working, press ^C to stop")
	wg.Wait()

	return c.Count()
}
```

```bash
 go test ../internal/concurrency -v -run="WorkingEndlesslyWithAGoodWayOut$" 
```

```bash
 go test ../internal/concurrency -v -run="WorkingEndlesslyWithAGoodWayOut$" -race 
```

## Result?

|                                                   Correctness                                                    |                                                   Consistency                                                    |                                                   Completeness                                                   |
|:----------------------------------------------------------------------------------------------------------------:|:----------------------------------------------------------------------------------------------------------------:|:----------------------------------------------------------------------------------------------------------------:|
| <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/question.svg" width="40"/> | <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/question.svg" width="40"/> | <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/question.svg" width="40"/> |

[Solution](example_7_solution.md)