# 7. Bonus question

[Safe counter](counter/safe.md)

```go
package concurrency

// NonStoppingGoRoutineCorrectShutdownBonus tiny change?
func NonStoppingGoRoutineCorrectShutdownBonus() int {
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

	fmt.Println("Working, press ^C to stop")
	wg.Wait()

	fmt.Printf("\nDid the go function shutdown gracefully? %v\n\n", gracefulShutdown)
	return c.Count()
}
```

```bash
 go test ../internal/concurrency -v -count=1 -run="NonStoppingGoRoutineCorrectShutdownBonus$" 
```

```bash
 go test ../internal/concurrency -v -count=1 -run="NonStoppingGoRoutineCorrectShutdownBonus$" -race 
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
    <td>Race conditions?</td>
    <td><img height="40" src="images/question.svg" width="40" alt="?"/></td> 
  </tr>
  <tr>
    <td>Error handling and gracefully shutdown?</td>
    <td><img height="40" src="images/question.svg" width="40" alt="?"/></td>
  </tr>
</tbody>
</table> 

[Back to README.md](../README.md)