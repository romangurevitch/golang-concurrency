# 5. Non-stopping go routines

[Safe counter](counter/safe.md)

```go
package concurrency

// NonStoppingGoRoutine is that a good idea?
func NonStoppingGoRoutine() int {
	c := &counter.SafeCounter{}
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		for {
			inlinePrint(c.Inc())
		}
	}()

	wg.Wait()
	return c.Count()
}
```

```bash
 go test ../internal/concurrency -v -count=1 -run="NonStoppingGoRoutine$" 
```

```bash
 go test ../internal/concurrency -v -count=1 -run="NonStoppingGoRoutine$" -race 
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
    <td><img height="40" src="images/no.png" width="40" alt="?"/></td>
    <td rowspan="3"><img height="320" src="https://media.giphy.com/media/lTrbUqQJCif7NfbXoo/giphy.gif" width="568" alt="?"/></td>
  </tr> 
  <tr>
    <td>Race conditions?</td>
    <td><img height="40" src="images/yes.png" width="40" alt="?"/></td> 
  </tr>
  <tr>
    <td>Error handling and gracefully shutdown?</td>
    <td><img height="40" src="images/no.png" width="40" alt="?"/></td>
  </tr>
</tbody>
</table> 

[Next example](example_6.md)
