# 1. Let's start with a simple example

[Simple counter](counter/simple.md)

```go 
package concurrency

// UnexpectedResult what did you expect? 
func UnexpectedResult() int {
	c := &counter.Simple{}

	go func() {
		for i := 0; i < 1000; i++ {
			c.Inc()
		}
	}()

	return c.Count()
}
```

```bash
 go test ../internal/concurrency -v -count=1 -run="UnexpectedResult$" 
```

```bash
 go test ../internal/concurrency -v -count=1 -run="UnexpectedResult$" -race 
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
    <td rowspan="3"><img height="360" src="https://media.giphy.com/media/xT0xeuOy2Fcl9vDGiA/giphy.gif" width="360" alt="?"/></td>
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

[Solution](example_1_solution.md)