# 2. Let's try and fix the issues

[Simple counter](counter/simple.md)

```go
package concurrency

// UnexpectedResultFix is it fixed?
func UnexpectedResultFix() int {
	c := &counter.SimpleCounter{}
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			c.Inc()
		}
	}()

	wg.Wait()
	return c.Count()
}
```

```bash
 clear; go test ../internal/concurrency -v -count=1 -run="UnexpectedResultFix$" 
```

```bash
 clear; go test ../internal/concurrency -v -count=1 -run="UnexpectedResultFix$" -race 
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
    <td rowspan="3"><img height="360" src="https://media.giphy.com/media/3onWp56oNIEHzEoPTE/giphy.gif" width="480" alt="?"/></td>
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

[Solution](example_2_solution.md)