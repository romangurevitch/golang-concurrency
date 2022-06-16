# 6. Let's handle shutdown gracefully?

[Safe counter](counter/safe.md)

```go
package concurrency

// NonStoppingGoRoutineWithShutdown is it good enough though?
func NonStoppingGoRoutineWithShutdown() int {
	c := &counter.SafeCounter{}
	gracefulShutdown := false

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		defer func() { gracefulShutdown = true }()

		for {
			inlinePrint(c.Inc())
		}
	}()

	fmt.Println("Working, press ^C to stop")
	<-sigs

	fmt.Printf("\nDid the go function shutdown gracefully? %v\n\n", gracefulShutdown)
	return c.Count()
}
```

```bash
 go test ../internal/concurrency -v -count=1 -run="NonStoppingGoRoutineWithShutdown$" 
```

```bash
 go test ../internal/concurrency -v -count=1 -run="NonStoppingGoRoutineWithShutdown$" -race 
```

## Result?

|                                                 Correctness                                                 |                                                 Consistency                                                 |                                                Completeness                                                 |
|:-----------------------------------------------------------------------------------------------------------:|:-----------------------------------------------------------------------------------------------------------:|:-----------------------------------------------------------------------------------------------------------:|
| <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/yes.png" width="40"/> | <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/yes.png" width="40"/> | <img height="40" src="/Users/RGurevitch/workspace/talk/golang-concurrency/docs/images/yes.png" width="40"/> |

[Next example](example_7.md)