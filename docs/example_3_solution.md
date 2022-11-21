# 3. Let's make a small change :)

[Simple counter](counter/simple.md)

```go
package concurrency

// LetsMakeASmallChange ohh no!
func LetsMakeASmallChange() int {
	c := &counter.SimpleCounter{}
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
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
 clear; go test ../internal/concurrency -v -count=1 -run="LetsMakeASmallChange$" 
```

```bash
 clear; go test ../internal/concurrency -v -count=1 -run="LetsMakeASmallChange$" -race 
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
    <td rowspan="3"><img height="360" src="https://media.giphy.com/media/lp1oGHyJHmSoqw0cld/giphy.gif" width="432" alt="?"/></td>
  </tr> 
  <tr>
    <td>No race conditions?</td>
    <td><img height="40" src="images/no.png" width="40" alt="?"/></td> 
  </tr>
  <tr>
    <td>Error handling and gracefully shutdown?</td>
    <td><img height="40" src="images/question.svg" width="40" alt="?"/></td>
  </tr>
</tbody>
</table> 

[Next example](example_4.md)