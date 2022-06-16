# 7. Let's handle shutdown gracefully, for real this time!

[Safe counter](counter/safe.md)

```go
package concurrency

// NonStoppingGoRoutineCorrectShutdown yes?
func NonStoppingGoRoutineCorrectShutdown() int {
	wg := sync.WaitGroup{}
	c := &counter.SafeCounter{}
	gracefulShutdown := false

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() { gracefulShutdown = true }()

		for {
			select {
			case <-sigs:
				return
			default:
				inlinePrint(c.Inc())
			}

		}
	}()

	fmt.Println("Working, press ^C to stop")
	wg.Wait()

	fmt.Printf("\nDid the go function shutdown gracefully? %v\n\n", gracefulShutdown)
	return c.Count()
}
```

```bash
 go test ../internal/concurrency -v -count=1 -run="NonStoppingGoRoutineCorrectShutdown$" 
```

```bash
 go test ../internal/concurrency -v -count=1 -run="NonStoppingGoRoutineCorrectShutdown$" -race 
```

## Result?

|                                                 Correctness                                                 |                                                 Consistency                                                 |                                                Completeness                                                 |
|:-----------------------------------------------------------------------------------------------------------:|:-----------------------------------------------------------------------------------------------------------:|:-----------------------------------------------------------------------------------------------------------:|
| <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/yes.png" width="40"/> | <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/yes.png" width="40"/> | <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/yes.png" width="40"/> |

[Back to README.md](../README.md)