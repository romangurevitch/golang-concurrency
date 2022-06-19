# 4. Finally, something works as expected - using safe counter

* [Simple counter](counter/simple.md)
* [Safe counter](counter/safe.md)

```go
package concurrency

// FinallySomethingWorksAsExpectedSafeCounter but is it?
func FinallySomethingWorksAsExpectedSafeCounter(count int) int {
	c := &counter.SafeCounter{}
	wg := sync.WaitGroup{}

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}

	wg.Wait()
	return c.Count()
}
```

```bash
 go test ../internal/concurrency -v -count=1 -run="FinallySomethingWorksAsExpectedSafeCounter$" 
```

```bash
 go test ../internal/concurrency -v -count=1 -run="FinallySomethingWorksAsExpectedSafeCounter$" -race 
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
    <td rowspan="3"><img height="360" src="https://media.giphy.com/media/f9Rrghj6TDckb5nZZR/giphy.gif" width="360" alt="?"/></td>
  </tr> 
  <tr>
    <td>Race conditions?</td>
    <td><img height="40" src="images/yes.png" width="40" alt="?"/></td> 
  </tr>
  <tr>
    <td>Error handling and gracefully shutdown?</td>
    <td><img height="40" src="images/question.svg" width="40" alt="?"/></td>
  </tr>
</tbody>
</table>

[Next example](example_5.md)