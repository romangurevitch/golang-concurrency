# 6. Let's handle shutdown gracefully?

[Safe counter](counter/safe.md)

```go
package concurrency

// NonStoppingGoRoutineWithShutdown is it good enough though?
func NonStoppingGoRoutineWithShutdown() (int, bool) {
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

	<-sigs
	return c.Count(), gracefulShutdown
}
```

```bash
 clear; go test ../internal/concurrency -v -count=1 -run="NonStoppingGoRoutineWithShutdown$" 
```

```bash
 clear; go test ../internal/concurrency -v -count=1 -run="NonStoppingGoRoutineWithShutdown$" -race 
```

<table>
<thead> 
  <tr> 
    <th colspan="3">Results?</th> 
  </tr>
</thead>
<tbody>
  <tr>
    <td>Correct result?</td>
    <td><img height="40" src="images/question.svg" width="40" alt="?"/></td>
    <td rowspan="3"><img height="320" src="https://media.giphy.com/media/48YKCwrp4Kt8I/giphy.gif" width="568" alt="?"/></td>
  </tr> 
  <tr>
    <td>No race conditions?</td>
    <td><img height="40" src="images/question.svg" width="40" alt="?"/></td> 
  </tr>
  <tr>
    <td>Error handling and gracefully shutdown?</td>
    <td><img height="40" src="images/question.svg" width="40" alt="?"/></td>
  </tr>
</tbody>
</table> 

[Solution](example_6_solution.md)