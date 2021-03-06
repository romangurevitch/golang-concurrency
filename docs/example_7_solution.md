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

<table>
<thead> 
  <tr> 
    <th colspan="3">Results?</th> 
  </tr>
</thead>
<tbody>
  <tr>
    <td>Correct result?</td>
    <td><img height="40" src="images/yes.png" width="40" alt="?"/></td>
    <td rowspan="3"><img height="320" src="https://media.giphy.com/media/3oxRmD9a5pLTOOLigM/giphy.gif" width="320" alt="?"/></td>
  </tr> 
  <tr>
    <td>No race conditions?</td>
    <td><img height="40" src="images/yes.png" width="40" alt="?"/></td> 
  </tr>
  <tr>
    <td>Error handling and gracefully shutdown?</td>
    <td><img height="40" src="images/yes.png" width="40" alt="?"/></td>
  </tr>
</tbody>
</table> 

[Bonus question](example_7_bonus.md)
