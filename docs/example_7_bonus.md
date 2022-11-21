# 7. Bonus question

[Safe counter](counter/safe.md)

```go
package concurrency

// NonStoppingGoRoutineCorrectShutdownBonus tiny change?
func NonStoppingGoRoutineCorrectShutdownBonus() (int, bool) {
	wg := sync.WaitGroup{}
	c := &counter.SafeCounter{}
	gracefulShutdown := false

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	wg.Add(1)
	go func() {
		defer func() { gracefulShutdown = true }() // <<< Tiny change
		defer wg.Done()                            // <<< Lets swap the defer commands

		for {
			select {
			case <-sigs:
				return
			default:
				inlinePrint(c.Inc())
			}

		}
	}()

	wg.Wait()
	return c.Count(), gracefulShutdown
}
```

```bash
 clear; go test ../internal/concurrency -v -count=1 -run="NonStoppingGoRoutineCorrectShutdownBonus$" 
```

```bash
 clear; go test ../internal/concurrency -v -count=1 -run="NonStoppingGoRoutineCorrectShutdownBonus$" -race 
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
    <td rowspan="3"><img height="320" src="https://media.giphy.com/media/U1TgwOffGUqxQYClV1/giphy.gif" width="568" alt="?"/></td>
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

[Back to README.md](../README.md)